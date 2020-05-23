package services

import (
	"io"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/bpetetot/leafer/db"
	"github.com/bpetetot/leafer/services/zip"
)

// MediaService gives access to media services
type MediaService struct {
	media db.MediaStore
}

// NewMediaService creates a media service instance
func NewMediaService(DB *gorm.DB) MediaService {
	return MediaService{
		media: db.NewMediaStore(DB),
	}
}

// Search search media corresponding to given query parameters
func (s *MediaService) Search(inputs db.SearchMediaInputs) (*[]db.Media, error) {
	return s.media.Search(inputs)
}

// Get get media info
func (s *MediaService) Get(id uint) (*db.Media, error) {
	return s.media.Get(id)
}

// StreamMediaPage return the media content
func (s *MediaService) StreamMediaPage(id uint, pageIndex int, w io.Writer) error {
	media, err := s.Get(id)
	if err != nil {
		return err
	}

	err = zip.ExtractImage(media.Path, pageIndex, w)
	if err != nil {
		return err
	}
	return nil
}

// MarkAsRead mark media as read
func (s *MediaService) MarkAsRead(id uint) error {
	viewedAt := time.Now()
	err := s.media.UpdateLastViewed(uint(id), &viewedAt)
	if err != nil {
		return err
	}
	return nil
}

// MarkAsUnread mark media as read
func (s *MediaService) MarkAsUnread(id uint) error {
	err := s.media.UpdateLastViewed(uint(id), nil)
	if err != nil {
		return err
	}
	return nil
}
