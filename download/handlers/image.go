package handlers

import (
	"context"
	errs "errors"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"math"
	"os"

	"github.com/disintegration/imaging"
	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/errors"
	"github.com/micro-blonde/file"
	f "github.com/micro-ginger/file/file/domain/file"
)

type imageHandler[T file.Model] struct {
	*base[T]

	config struct {
		SupportedSizes []uint64
	}
}

func NewImage[T file.Model](registry registry.Registry) Handler[T] {
	h := &imageHandler[T]{
		base: newBase[T](),
	}
	if err := registry.Unmarshal(&h.config); err != nil {
		panic(err)
	}
	return h
}

func (h *imageHandler[T]) Resize(ctx context.Context, refFile *f.File[T],
	request Request) (resizedAbsPath string, err errors.Error) {
	if request.GetMaxSize() == 0 {
		return "", nil
	}
	var maxSize uint64 = 0
	for _, size := range h.config.SupportedSizes {
		if maxSize == 0 || (size <= request.GetMaxSize() && maxSize < size) {
			maxSize = size
		}
	}
	//
	resizedAbsPath = h.storage.GetTempAbsPath(refFile, fmt.Sprint(maxSize))
	f, oErr := os.Open(resizedAbsPath)
	if oErr != nil {
		if !errs.Is(oErr, os.ErrNotExist) {
			return "", errors.New(oErr).
				WithTrace("os.Open")
		}
		// load original file
		origFile, fErr := os.Open(h.storage.GetAbsPath(refFile))
		if fErr != nil {
			return "", errors.New(fErr).
				WithTrace("original>os.Open")
		}
		defer func() {
			closeErr := origFile.Close()
			if err != nil {
				err = errors.New(closeErr).
					WithTrace("original>file.Close")
			}
		}()
		imageConf, _, dErr := image.DecodeConfig(origFile)
		if dErr != nil {
			return "", errors.New(dErr).
				WithTrace("curFile>image.DecodeConfig")
		}
		max := request.GetMaxSize()
		resizeRatio := math.Max(
			math.Min(float64(max), float64(imageConf.Width))/float64(imageConf.Width),
			math.Min(float64(max), float64(imageConf.Height))/float64(imageConf.Height),
		)
		if resizeRatio >= 1 {
			return "", nil
		}
		// resize
		reader, newWidth, newHeight, err := h.
			ensureImageMaxSize(origFile, max)
		if err != nil {
			return "", err.
				WithTrace("ensureImageMaxSize")
		}
		if newWidth <= 0 || newHeight <= 0 {
			return "", nil
		}
		// store
		dir := h.storage.GetTempRelativeDirPath(refFile.Key)
		req := &file.StoreRequest{
			Reader:    reader,
			Category:  refFile.Category,
			Type:      refFile.Type,
			Name:      refFile.Name,
			Extension: refFile.Extension,
			AccountId: refFile.AccountId,
			BaseDir:   &dir,
		}
		refFile.Key = h.storage.GetTempFileName(refFile, fmt.Sprint(maxSize))
		result, err := h.storage.SaveFile(ctx, req, refFile)
		if err != nil {
			return "", err.
				WithTrace("sotorage.Store")
		}
		resizedAbsPath = result.AbsPath
		return resizedAbsPath, nil
	}
	f.Close()
	return resizedAbsPath, nil
}

func (h *imageHandler[T]) getValidMaxSize(max uint64) uint64 {
	if h.config.SupportedSizes != nil {
		var lastSize uint64
		for _, size := range h.config.SupportedSizes {
			lastSize = size
			if lastSize >= max {
				break
			}
		}
		max = lastSize
	}
	return max
}

// ensureImageMaxSize checks the given image file and resizes image to maximum
// configured sizes existing in the config matching with given max param
func (h *imageHandler[T]) ensureImageMaxSize(file io.ReadSeeker,
	max uint64) (out io.ReadSeeker, width, height uint64, err errors.Error) {
	max = h.getValidMaxSize(max)

	if _, sErr := file.Seek(0, io.SeekStart); sErr != nil {
		err = errors.New(sErr).
			WithTrace("file.Seek")
		return
	}
	if max <= 0 {
		out = file
		return
	}
	imageConf, _, dErr := image.DecodeConfig(file)
	if dErr != nil {
		err = errors.New(dErr).
			WithTrace("image.DecodeConfig")
		return
	}
	resizeRatio := math.Max(
		math.Min(float64(max), float64(imageConf.Width))/float64(imageConf.Width),
		math.Min(float64(max), float64(imageConf.Height))/float64(imageConf.Height),
	)
	if resizeRatio >= 1 {
		out = file
		return
	}
	width = uint64(float64(imageConf.Width) * resizeRatio)
	height = uint64(float64(imageConf.Height) * resizeRatio)
	file.Seek(0, io.SeekStart)
	originalImage, _, dErr := image.Decode(file)
	if dErr != nil {
		err = errors.New(dErr).
			WithTrace("image.Decode")
		return
	}
	newImage := imaging.Resize(originalImage,
		int(width), int(height), imaging.Lanczos)
	var rws ReadWriterSeeker
	rws.InitializeRWS()
	eErr := jpeg.Encode(&rws, newImage, nil)
	if eErr != nil {
		err = errors.New(eErr).
			WithTrace("jpeg.Encode")
		return
	}
	out = rws.GetReadSeeker()
	_, sErr := out.Seek(0, io.SeekStart)
	if sErr != nil {
		err = errors.New(sErr).
			WithTrace("out.Seek")
		return
	}
	return
}
