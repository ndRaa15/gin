package repository

import (
	"context"

	"gin/internal/domain/entity"
	"gin/internal/infrastructure/mysql"
)

type StudioRepository struct {
	db *mysql.DB
}

func NewStudioRepository(db *mysql.DB) *StudioRepository {
	return &StudioRepository{
		db: db,
	}
}

func (sr *StudioRepository) GetAll(ctx context.Context) ([]*entity.Studio, error) {
	var studios []*entity.Studio
	if err := sr.db.Debug().WithContext(ctx).Preload("EndTime.Time").Preload("StartTime.Time").Find(&studios).Error; err != nil {
		return nil, err
	}
	return studios, nil
}

func (sr *StudioRepository) GetByID(ctx context.Context, id uint) (*entity.Studio, error) {
	var studio entity.Studio
	if err := sr.db.Debug().WithContext(ctx).Preload("EndTime.Time").Preload("StartTime.Time").Where("id = ?", id).First(&studio).Error; err != nil {
		return nil, err
	}
	return &studio, nil
}

func (sr *StudioRepository) Update(ctx context.Context, studio *entity.Studio) (*entity.Studio, error) {
	if err := sr.db.Debug().WithContext(ctx).Save(studio).Error; err != nil {
		return nil, err
	}
	return studio, nil
}

func (sr *StudioRepository) UpdateStartTime(ctx context.Context, studioID, timeID uint, status bool) error {
	if err := sr.db.Debug().WithContext(ctx).Model(entity.StartTime{}).Where("studio_id = ? AND time_id = ?", studioID, timeID).Update("is_available", status).Error; err != nil {
		return err
	}
	return nil
}

func (sr *StudioRepository) UpdateEndTime(ctx context.Context, studioID, timeID uint, status bool) error {
	if err := sr.db.Debug().WithContext(ctx).Model(entity.EndTime{}).Where("studio_id = ? AND time_id = ?", studioID, timeID).Update("is_available", status).Error; err != nil {
		return err
	}
	return nil
}

func (sr *StudioRepository) RentStudio(ctx context.Context, rentStudio *entity.RentStudio) (*entity.RentStudio, error) {
	if err := sr.db.Debug().WithContext(ctx).Create(&rentStudio).Error; err != nil {
		return nil, err
	}
	return rentStudio, nil
}

func SeedStudio(db *mysql.DB) error {
	var studios []*entity.Studio

	studio1 := entity.Studio{
		ID:           1,
		Name:         "Tom Studio Malang",
		Address:      "Jl. Raya Tlogomas No.246, Malang",
		Description:  "",
		PricePerHour: 100_000,
		OpenHour:     "08:00",
		Phone:        "+6282145278752",
		Rating:       4.5,
		Photo:        "https.drive.google.com/uc?id=1xr7lzsYF-dLdEq-VnNqrNTTxS3TBL3OD",
	}

	var startTime1 []entity.StartTime
	for i := 7; i <= 18; i++ {
		entity := entity.StartTime{
			StudioID:    1,
			TimeID:      uint(i),
			IsAvailable: true,
		}
		startTime1 = append(startTime1, entity)
	}

	var endTime1 []entity.EndTime
	for i := 8; i <= 19; i++ {
		entity := entity.EndTime{
			StudioID:    1,
			TimeID:      uint(i),
			IsAvailable: true,
		}
		endTime1 = append(endTime1, entity)
	}

	studio2 := entity.Studio{
		ID:           2,
		Name:         "Corner Music Studio",
		Address:      "123 Main Street, Malang",
		Description:  "Corner Music Studio adalah tempat yang sempurna untuk menggali bakat musik Anda. Dengan fasilitas modern dan instruktur berpengalaman, kami membantu Anda mencapai potensi musik Anda.",
		PricePerHour: 80_000,
		OpenHour:     "10:00",
		Phone:        "+6282145278752",
		Rating:       4.3,
		Photo:        "https.drive.google.com/uc?id=198EprxIVVYRK23-maBYPow05mY_QdEZw",
	}

	var startTime2 []entity.StartTime
	for i := 8; i <= 18; i++ {
		entity := entity.StartTime{
			StudioID: 2,
			TimeID:   uint(i),
		}
		startTime2 = append(startTime2, entity)
	}

	var endTime2 []entity.EndTime
	for i := 9; i <= 19; i++ {
		entity := entity.EndTime{
			StudioID: 2,
			TimeID:   uint(i),
		}
		endTime2 = append(endTime2, entity)
	}

	studio3 := entity.Studio{
		ID:           3,
		Name:         "Virtuoso Music Studio",
		Address:      "456 Elm Avenue, Malang",
		Description:  "Virtuoso Music Studio adalah tempat di mana bakat bertemu permainan yang luar biasa. Dengan lingkungan yang kreatif dan pelatih yang berdedikasi, Anda akan mencapai prestasi tertinggi dalam musik.",
		PricePerHour: 120_000,
		OpenHour:     "09:00",
		Phone:        "+6282145278752",
		Rating:       4.7,
		Photo:        "https.drive.google.com/uc?id=1rr63uvZ_9rVAJLFRn4BZdq11jz3XQKMA",
	}

	var startTime3 []entity.StartTime
	for i := 7; i <= 15; i++ {
		entity := entity.StartTime{
			StudioID: 3,
			TimeID:   uint(i),
		}
		startTime3 = append(startTime3, entity)
	}

	var endTime3 []entity.EndTime
	for i := 8; i <= 16; i++ {
		entity := entity.EndTime{
			StudioID: 3,
			TimeID:   uint(i),
		}
		endTime3 = append(endTime3, entity)
	}

	studio4 := entity.Studio{
		ID:           4,
		Name:         "King and Queen Music Studio",
		Address:      "789 Oak Street, Malang",
		Description:  "King and Queen Music Studio adalah tempat yang dihormati oleh semua musisi. Kami menawarkan peralatan berkualitas tinggi dan suasana yang inspiratif untuk menghasilkan musik yang luar biasa.",
		PricePerHour: 150_000,
		OpenHour:     "11:00",
		Phone:        "+6282145278752",
		Rating:       4.1,
		Photo:        "https.drive.google.com/uc?id=1Q8f2qiHsDEtDCHGWOkJ8bMPeuYG7a3kW",
	}

	var startTime4 []entity.StartTime
	for i := 7; i <= 18; i++ {
		entity := entity.StartTime{
			StudioID: 4,
			TimeID:   uint(i),
		}
		startTime4 = append(startTime4, entity)
	}

	var endTime4 []entity.EndTime
	for i := 8; i <= 19; i++ {
		entity := entity.EndTime{
			StudioID: 4,
			TimeID:   uint(i),
		}
		endTime4 = append(endTime4, entity)
	}

	studio5 := entity.Studio{
		ID:           5,
		Name:         "Rumah Music Mendut",
		Address:      "101 Pine Road, Malang",
		Description:  "Rumah Music Mendut adalah tempat yang hangat dan ramah untuk mengejar hasrat musik Anda. Bergabunglah dengan komunitas musik kami dan nikmati perjalanan musik Anda bersama kami.",
		PricePerHour: 70_000,
		OpenHour:     "12:00",
		Phone:        "+6282145278752",
		Rating:       4.3,
		Photo:        "https.drive.google.com/uc?id=1MDH72RXDkCBXLV4KX3FldRiziL6ivhov",
	}

	var startTime5 []entity.StartTime
	for i := 7; i <= 18; i++ {
		entity := entity.StartTime{
			StudioID: 5,
			TimeID:   uint(i),
		}
		startTime5 = append(startTime5, entity)
	}

	var endTime5 []entity.EndTime
	for i := 8; i <= 19; i++ {
		entity := entity.EndTime{
			StudioID: 5,
			TimeID:   uint(i),
		}
		endTime5 = append(endTime5, entity)
	}

	studio6 := entity.Studio{
		ID:           6,
		Name:         "Aria Music Studio 651",
		Address:      "651 Cedar Lane, Malang",
		Description:  "Aria Music Studio 651 adalah tempat untuk mengejar harmoni. Kami memiliki beragam instrumen dan pelajaran untuk memenuhi semua jenis minat musik Anda.",
		PricePerHour: 110_000,
		OpenHour:     "08:30",
		Phone:        "+6282145278752",
		Rating:       4.7,
		Photo:        "https.drive.google.com/uc?id=18sWFeg35aSOsfPIRAfZOjLgJXJPxpBv_",
	}

	var startTime6 []entity.StartTime
	for i := 7; i <= 18; i++ {
		entity := entity.StartTime{
			StudioID: 6,
			TimeID:   uint(i),
		}
		startTime6 = append(startTime6, entity)
	}

	var endTime6 []entity.EndTime
	for i := 8; i <= 19; i++ {
		entity := entity.EndTime{
			StudioID: 6,
			TimeID:   uint(i),
		}
		endTime6 = append(endTime6, entity)
	}

	studio7 := entity.Studio{
		ID:           7,
		Name:         "AA Music Studio",
		Address:      "222 Willow Avenue, Malang",
		Description:  "AA Music Studio adalah tempat yang menyenangkan untuk belajar dan berkembang dalam musik. Mari bergabung dengan kami dan nikmati perjalanan musik yang menyenangkan.",
		PricePerHour: 90_000,
		OpenHour:     "10:30",
		Phone:        "+6282145278752",
		Rating:       4.9,
		Photo:        "https.drive.google.com/uc?id=1VtW1LhASD27mA9bH73-a8SS3gZUnz4GD",
	}

	var startTime7 []entity.StartTime
	for i := 7; i <= 18; i++ {
		entity := entity.StartTime{
			StudioID: 7,
			TimeID:   uint(i),
		}
		startTime7 = append(startTime7, entity)
	}

	var endTime7 []entity.EndTime
	for i := 8; i <= 19; i++ {
		entity := entity.EndTime{
			StudioID: 7,
			TimeID:   uint(i),
		}
		endTime7 = append(endTime7, entity)
	}

	studio8 := entity.Studio{
		ID:           8,
		Name:         "Bluesky",
		Address:      "333 Birch Street, Malang",
		Description:  "Bluesky adalah studio musik dengan nuansa yang santai dan pemandangan yang indah. Rasakan ketenangan dan inspirasi dalam musik di sini.",
		PricePerHour: 75_000,
		OpenHour:     "09:30",
		Phone:        "+6282145278752",
		Rating:       4.1,
		Photo:        "https.drive.google.com/uc?id=1ubH0g4bxdWCWBvarXkrWj7olGIVR8vQY",
	}

	var startTime8 []entity.StartTime
	for i := 7; i <= 18; i++ {
		entity := entity.StartTime{
			StudioID: 8,
			TimeID:   uint(i),
		}
		startTime8 = append(startTime8, entity)
	}

	var endTime8 []entity.EndTime
	for i := 8; i <= 19; i++ {
		entity := entity.EndTime{
			StudioID: 8,
			TimeID:   uint(i),
		}
		endTime8 = append(endTime8, entity)
	}

	studio9 := entity.Studio{
		ID:           9,
		Name:         "Md Music Studio",
		Address:      "444 Maple Drive, Malang",
		Description:  "Md Music Studio adalah tempat di mana musik dan inovasi bertemu. Kami menawarkan berbagai kursus musik untuk memenuhi kebutuhan Anda.",
		PricePerHour: 100_000,
		OpenHour:     "11:30",
		Phone:        "+6282145278752",
		Rating:       4.4,
		Photo:        "https.drive.google.com/uc?id=1YC-frk5Zcb3nZ5KTdVNgEy6oWSRvawXn",
	}

	var startTime9 []entity.StartTime
	for i := 7; i <= 18; i++ {
		entity := entity.StartTime{
			StudioID: 9,
			TimeID:   uint(i),
		}
		startTime9 = append(startTime9, entity)
	}

	var endTime9 []entity.EndTime
	for i := 8; i <= 19; i++ {
		entity := entity.EndTime{
			StudioID: 9,
			TimeID:   uint(i),
		}
		endTime9 = append(endTime9, entity)
	}

	studio10 := entity.Studio{
		ID:           10,
		Name:         "Virtuoso",
		Address:      "555 Oakwood Road, Sidoarjo",
		Description:  "Virtuoso adalah studio musik yang menghormati keunggulan. Bergabunglah dengan kami dan kembangkan bakat musik Anda dengan instruktur terbaik.",
		PricePerHour: 130_000,
		OpenHour:     "08:00",
		Phone:        "+6282145278752",
		Rating:       4.1,
		Photo:        "https.drive.google.com/uc?id=1H2lrUBjyg9hlKX4Iwb9-cZfLk4G1leNi",
	}

	var startTime10 []entity.StartTime
	for i := 7; i <= 18; i++ {
		entity := entity.StartTime{
			StudioID: 10,
			TimeID:   uint(i),
		}
		startTime10 = append(startTime10, entity)
	}

	var endTime10 []entity.EndTime
	for i := 8; i <= 19; i++ {
		entity := entity.EndTime{
			StudioID: 10,
			TimeID:   uint(i),
		}
		endTime10 = append(endTime10, entity)
	}

	studios = append(studios, &studio1, &studio2, &studio3, &studio4, &studio5, &studio6, &studio7, &studio8, &studio9, &studio10)
	var startTime [][]entity.StartTime
	startTime = append(startTime, startTime1, startTime2, startTime3, startTime4, startTime5, startTime6, startTime7, startTime8, startTime9, startTime10)
	var endTime [][]entity.EndTime
	endTime = append(endTime, endTime1, endTime2, endTime3, endTime4, endTime5, endTime6, endTime7, endTime8, endTime9, endTime10)

	for i := 0; i < len(studios); i++ {
		if err := db.Create(studios[i]).Error; err != nil {
			return err
		}
		if err := db.Model(entity.StartTime{}).Create(startTime[i]).Error; err != nil {
			return err
		}
		if err := db.Model(entity.EndTime{}).Create(endTime[i]).Error; err != nil {
			return err
		}
	}
	return nil
}
