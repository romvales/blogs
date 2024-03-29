package store

import (
	"errors"
	"log"

	"github.com/romv7/blogs/internal/pb"
	"github.com/romv7/blogs/internal/storage"
	sqlStore "github.com/romv7/blogs/internal/store/sql"
	sqlModels "github.com/romv7/blogs/internal/store/sql/models"
	"github.com/romv7/blogs/internal/utils/author"

	"gorm.io/gorm"
)

type User struct {
	t        StoreType
	s        storage.StorageDriverType
	sqlModel *sqlModels.User
}

type UserStore struct {
	t StoreType
	s storage.StorageDriverType
}

func NewUserStore(t StoreType) *UserStore {
	return &UserStore{t, storage.Plain}
}

func (s *UserStore) GetMainStore() (S any) {
	switch s.t {
	case SqlStore:
		S = sqlStore.Store()
	default:
		log.Panic(ErrInvalidStore)
	}

	return
}

func (s *UserStore) NewUser(u *pb.User) (out *User) {
	out = &User{}

	switch s.t {
	case SqlStore:
		out.t = SqlStore
		out.sqlModel = sqlModels.NewUser(u)
	default:
		log.Panic(ErrInvalidStore)
	}

	// Set the storage type
	out.s = s.s

	// Populates the metadata of an author user using the stored metadata
	// from a storage.
	pbuser := out.Proto()

	if pbuser.Type == pb.User_T_AUTHOR {
		ah := author.NewAuthorHelper(pbuser, storage.Plain)
		metadata := ah.GetAuthorMetadata()

		if metadata == nil {
			return
		}

		pbuser.Bio = metadata.Bio
		pbuser.AltName = metadata.AltName

		for plat, links := range metadata.SocialLinks {
			pbuser.SocialLinks[string(plat)] = &pb.SocialLinks{Data: links}
		}
	}

	return
}

func (s *UserStore) Save(u *User) (err error) {

	switch s.t {
	case SqlStore:
		db := sqlStore.Store()
		res := db.Save(u.sqlModel)
		err = res.Error
	default:
		log.Panic(ErrInvalidStore)
	}

	// Saves the user metadata (if any) to the storage.
	pbuser := u.Proto()

	if pbuser.Type == pb.User_T_AUTHOR {
		ah := author.NewAuthorHelper(pbuser, storage.Plain)
		ah.SaveAuthorMetadata()
	}

	return
}

func (s *UserStore) Delete(u *User) (err error) {
	user := u.Proto()

	switch s.t {
	case SqlStore:
		db := sqlStore.Store()

		if res := db.Where("uuid = ?", user.Uuid).Delete(u.sqlModel); res.Error != nil {
			return res.Error
		}
	default:
		log.Panic(ErrInvalidStore)
	}

	if user.Type == pb.User_T_AUTHOR {
		ah := author.NewAuthorHelper(user, storage.Plain)
		err = ah.DeleteAuthorMetadata()
	}

	return
}

// PROBLEM: Sometimes GetById is returning record not found even though the record exist.

// RESOLVED:
//
//	Turns out that this function is not the problem, the reason why this method
//	is returning a record not found is that the unit test were tightly interconnected
//	as a result, when one test deleted a user and the other tries to GetById the record it
//	is apparent that it should return a not found error because the other test already deleted
//	the record we are trying to lookup with GetById.

// TODO: Add a documentation to this method.
func (s *UserStore) GetById(id uint64) (out *User, err error) {
	out = &User{}

	switch s.t {
	case SqlStore:
		db := sqlStore.Store()
		out.t = s.t
		out.sqlModel = &sqlModels.User{ID: id}

		if res := db.First(out.sqlModel); errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, res.Error
		}

		if out.sqlModel.ID != id {
			return nil, gorm.ErrRecordNotFound
		}
	default:
		log.Panic(ErrInvalidStore)
	}

	return
}

// TODO: Add a documentation to this method.
func (s *UserStore) GetByUuid(uuid string) (out *User, err error) {
	out = &User{}

	switch s.t {
	case SqlStore:
		db := sqlStore.Store()
		out.t = s.t
		out.sqlModel = &sqlModels.User{}

		if res := db.Where("uuid = ?", uuid).Limit(1).Find(out.sqlModel); errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, res.Error
		}

		if out.sqlModel.Uuid != uuid {
			return nil, gorm.ErrRecordNotFound
		}
	default:
		log.Panic(ErrInvalidStore)
	}

	return
}
