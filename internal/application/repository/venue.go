package repository

import (
	"context"

	"gin/internal/domain/entity"
	"gin/internal/domain/repository"
	"gin/internal/infrastructure/mysql"
	"github.com/gofrs/uuid"
)

type VenueRepository struct {
	db *mysql.DB
}

func NewVenueRepository(db *mysql.DB) repository.VenueRepositoryImpl {
	return &VenueRepository{db}
}

func (vr *VenueRepository) GetAll(ctx context.Context) ([]*entity.Venue, error) {
	var venues []*entity.Venue
	if err := vr.db.Debug().WithContext(ctx).Preload("VenueDays.Day").Find(&venues).Error; err != nil {
		return nil, err
	}
	return venues, nil
}

func (vr *VenueRepository) GetByID(ctx context.Context, id uint) (*entity.Venue, error) {
	var venue entity.Venue
	if err := vr.db.Debug().WithContext(ctx).Preload("VenueDays.Day").Where("id = ?", id).First(&venue).Error; err != nil {
		return nil, err
	}
	return &venue, nil
}

func (vr *VenueRepository) GetVenueDayByID(ctx context.Context, venueDayID uint) (*entity.VenueDay, error) {
	var venueDay entity.VenueDay
	if err := vr.db.Debug().WithContext(ctx).Preload("Day").Where("id = ?", venueDayID).First(&venueDay).Error; err != nil {
		return nil, err
	}
	return &venueDay, nil
}

func (vr *VenueRepository) CreateApplyVenue(ctx context.Context, applyVenue *entity.ApplyVenue) (*entity.ApplyVenue, error) {
	if err := vr.db.Debug().WithContext(ctx).Create(&applyVenue).Preload("VenueDay").Error; err != nil {
		return nil, err
	}
	return applyVenue, nil
}

func (vr *VenueRepository) GetApplyVenueByUserID(ctx context.Context, userID uuid.UUID) ([]*entity.ApplyVenue, error) {
	var applyVenues []*entity.ApplyVenue
	if err := vr.db.Debug().WithContext(ctx).Where("user_id = ?", userID).Find(&applyVenues).Error; err != nil {
		return nil, err
	}
	return applyVenues, nil
}

func SeedVenue(db *mysql.DB) error {
	var venues []entity.Venue
	var venueDays []entity.VenueDay

	venue1 := entity.Venue{
		ID:          1,
		Name:        "Nakoa Soekarno Hatta Malang",
		Address:     "Jl. Soalawas, No. 1",
		Description: "Kopi Soa adalah tempat nongkrong yang asik",
		Photo:       "https://drive.google.com/uc?id=1XNFbVUBYXf97zDq3Kmw4FxkK5yacRWoS",
	}

	venue2 := entity.Venue{
		ID:          2,
		Name:        "Inti Kopi",
		Address:     "Jl. Pahlawan, No. 10",
		Description: "Inti Kopi adalah tempat nongkrong yang cozy",
		Photo:       "https://drive.google.com/uc?id=1tFcvrs29gV2hMAPvDS8HR2Ialc07c4Sg",
	}

	venue3 := entity.Venue{
		ID:          3,
		Name:        "11:12",
		Address:     "Jl. Nyaman, No. 5",
		Description: "11:12 adalah tempat makan enak dan murah",
		Photo:       "https://drive.google.com/uc?id=1j2uk6AliqhK-qsuCtKgKznUEYGlX8pVh",
	}

	venue4 := entity.Venue{
		ID:          4,
		Name:        "Stako Coffee Malang",
		Address:     "Jl. Kedai Kopi, No. 7",
		Description: "Stako Coffee adalah tempat yang menghipnotis Anda dengan aroma kopi segar, suasana yang hangat, dan kenyamanan yang tak tertandingi. Nikmati momen istimewa dengan secangkir kopi pilihan Anda di Stako Coffee Malang.",
		Photo:       "https://drive.google.com/uc?id=1LA28_eCKK20erD1Nye5AMk09y3dlSAvT",
	}

	venue5 := entity.Venue{
		ID:          5,
		Name:        "Pesenkopi+",
		Address:     "Jl. Aroma, No. 3",
		Description: "Pesenkopi+ adalah destinasi kopi yang memukau dengan berbagai varian kopi dari seluruh penjuru dunia. Setiap tegukan adalah petualangan rasa yang tak terlupakan, jadikan setiap kunjungan Anda di Pesenkopi+ sebagai perjalanan pengejaran cita rasa.",
		Photo:       "https://drive.google.com/uc?id=1fHV-ip0HVneIgINvxWRxgFyfNAVy0Hgb",
	}

	venue6 := entity.Venue{
		ID:          6,
		Name:        "Lavayette",
		Address:     "Jl. Lavender, No. 12",
		Description: "Lavayette adalah tempat eksklusif yang mengundang Anda ke dalam dunia elegan kopi dan teh. Dengan suasana yang tenang dan dekorasi yang mewah, Lavayette adalah destinasi sempurna untuk merayakan momen-momen istimewa Anda.",
		Photo:       "https://drive.google.com/uc?id=1ZAeeMY50mJ2tttn8MhQTgP94v05ubunV",
	}

	venue7 := entity.Venue{
		ID:          7,
		Name:        "Kopi Dari Hati",
		Address:     "Jl. Hati Indah, No. 8",
		Description: "Kopi Dari Hati adalah surga bagi pecinta kopi sejati. Di sini, kami meracik kopi dengan cinta dan dedikasi yang mendalam. Setiap cangkir adalah karya seni, dan kami mengundang Anda untuk menjelajahi keindahan kopi bersama kami.",
		Photo:       "https://drive.google.com/uc?id=1Nmm6cLPzJ242PIQxhluqUC0A_8UNSTCU",
	}

	venue8 := entity.Venue{
		ID:          8,
		Name:        "Warkop Brewok",
		Address:     "Jl. Jenggot, No. 9",
		Description: "Warkop Brewok adalah tempat yang penuh karakter dan kehangatan. Di sini, kami menghidangkan kopi berkualitas tinggi sambil berbagi cerita dan tawa. Bergabunglah dengan kami di Warkop Brewok untuk pengalaman yang menyenangkan dan tak terlupakan.",
		Photo:       "https://drive.google.com/uc?id=1BOYKJU5uMrYQOVbM6y2Ids_v9pwUZGx3",
	}

	venue9 := entity.Venue{
		ID:          9,
		Name:        "Malang Creative Area",
		Address:     "Jl. Kreativitas, No. 11",
		Description: "Malang Creative Area adalah pusat seni dan kreativitas yang brimming dengan inspirasi. Nikmati seni lokal, pertunjukan, dan acara-acara unik di tempat ini. Bergabunglah dengan kami di Malang Creative Area dan jadilah bagian dari komunitas kreatif kami.",
		Photo:       "https://drive.google.com/uc?id=1GthHGp6m3VW-FOPQJWJBOHtp0aR62v_x",
	}

	venue10 := entity.Venue{
		ID:          10,
		Name:        "Fath",
		Address:     "Jl. Kesejahteraan, No. 6",
		Description: "Fath adalah destinasi kuliner yang memanjakan lidah Anda dengan hidangan lezat dari seluruh dunia. Dengan pelayanan ramah dan suasana yang hangat, Fath adalah tempat yang sempurna untuk menjalani perjalanan kuliner Anda.",
		Photo:       "https://drive.google.com/uc?id=1Pkf1HwhTnWSXn-cJYGnyXQZUQDgRtM9h",
	}

	venues = append(venues, venue1, venue2, venue3, venue4, venue5, venue6, venue7, venue8, venue9, venue10)

	for _, venue := range venues {
		if err := db.Create(&venue).Error; err != nil {
			return err
		}
	}

	venue1Days := entity.VenueDay{
		VenueID:   1,
		DayID:     1,
		Salary:    200_000,
		StartTime: "10.00",
		EndTime:   "12.00",
	}

	venue2Days := entity.VenueDay{
		VenueID:   1,
		DayID:     1,
		Salary:    200_000,
		StartTime: "13.00",
		EndTime:   "15.00",
	}

	venue3Days := entity.VenueDay{
		VenueID:   1,
		DayID:     1,
		Salary:    200_000,
		StartTime: "15.00",
		EndTime:   "17.00",
	}

	venue4Days := entity.VenueDay{
		VenueID:   1,
		DayID:     1,
		Salary:    200_000,
		StartTime: "19.00",
		EndTime:   "21.00",
	}

	venue5Days := entity.VenueDay{
		VenueID:   1,
		DayID:     2,
		Salary:    200_000,
		StartTime: "10.00",
		EndTime:   "12.00",
	}

	venue6Days := entity.VenueDay{
		VenueID:   1,
		DayID:     2,
		Salary:    200_000,
		StartTime: "13.00",
		EndTime:   "15.00",
	}

	venue7Days := entity.VenueDay{
		VenueID:   1,
		DayID:     2,
		Salary:    200_000,
		StartTime: "15.00",
		EndTime:   "17.00",
	}

	venue8Days := entity.VenueDay{
		VenueID:   1,
		DayID:     2,
		Salary:    200_000,
		StartTime: "19.00",
		EndTime:   "21.00",
	}

	venue9Days := entity.VenueDay{
		VenueID:   2,
		DayID:     2,
		Salary:    150_000,
		StartTime: "10.00",
		EndTime:   "12.00",
	}

	venue10Days := entity.VenueDay{
		VenueID:   2,
		DayID:     2,
		Salary:    150_000,
		StartTime: "13.00",
		EndTime:   "15.00",
	}

	venue11Days := entity.VenueDay{
		VenueID:   2,
		DayID:     2,
		Salary:    150_000,
		StartTime: "15.00",
		EndTime:   "17.00",
	}

	venue12Days := entity.VenueDay{
		VenueID:   2,
		DayID:     2,
		Salary:    150_000,
		StartTime: "19.00",
		EndTime:   "21.00",
	}

	venue13Days := entity.VenueDay{
		VenueID:   3,
		DayID:     3,
		Salary:    250_000,
		StartTime: "10.00",
		EndTime:   "12.00",
	}

	venue14Days := entity.VenueDay{
		VenueID:   3,
		DayID:     3,
		Salary:    250_000,
		StartTime: "13.00",
		EndTime:   "15.00",
	}

	venue15Days := entity.VenueDay{
		VenueID:   3,
		DayID:     3,
		Salary:    250_000,
		StartTime: "15.00",
		EndTime:   "17.00",
	}

	venue16Days := entity.VenueDay{
		VenueID:   3,
		DayID:     3,
		Salary:    250_000,
		StartTime: "19.00",
		EndTime:   "21.00",
	}

	venue17Days := entity.VenueDay{
		VenueID:   4,
		DayID:     4,
		Salary:    100_000,
		StartTime: "10.00",
		EndTime:   "12.00",
	}

	venue18Days := entity.VenueDay{
		VenueID:   4,
		DayID:     4,
		Salary:    100_000,
		StartTime: "13.00",
		EndTime:   "15.00",
	}

	venue19Days := entity.VenueDay{
		VenueID:   4,
		DayID:     4,
		Salary:    100_000,
		StartTime: "15.00",
		EndTime:   "17.00",
	}

	venue20Days := entity.VenueDay{
		VenueID:   4,
		DayID:     4,
		Salary:    100_000,
		StartTime: "19.00",
		EndTime:   "21.00",
	}

	venue21Days := entity.VenueDay{
		VenueID:   5,
		DayID:     5,
		Salary:    100_000,
		StartTime: "10.00",
		EndTime:   "12.00",
	}

	venue22Days := entity.VenueDay{
		VenueID:   5,
		DayID:     5,
		Salary:    100_000,
		StartTime: "13.00",
		EndTime:   "15.00",
	}

	venue23Days := entity.VenueDay{
		VenueID:   5,
		DayID:     5,
		Salary:    100_000,
		StartTime: "15.00",
		EndTime:   "17.00",
	}

	venue24Days := entity.VenueDay{
		VenueID:   5,
		DayID:     5,
		Salary:    100_000,
		StartTime: "19.00",
		EndTime:   "21.00",
	}

	venue25Days := entity.VenueDay{
		VenueID:   6,
		DayID:     6,
		Salary:    250_000,
		StartTime: "10.00",
		EndTime:   "12.00",
	}

	venue26Days := entity.VenueDay{
		VenueID:   6,
		DayID:     6,
		Salary:    250_000,
		StartTime: "13.00",
		EndTime:   "15.00",
	}

	venue27Days := entity.VenueDay{
		VenueID:   6,
		DayID:     6,
		Salary:    250_000,
		StartTime: "15.00",
		EndTime:   "17.00",
	}

	venue28Days := entity.VenueDay{
		VenueID:   6,
		DayID:     6,
		Salary:    250_000,
		StartTime: "19.00",
		EndTime:   "21.00",
	}

	venue29Days := entity.VenueDay{
		VenueID:   7,
		DayID:     7,
		Salary:    170_000,
		StartTime: "10.00",
		EndTime:   "12.00",
	}

	venue30Days := entity.VenueDay{
		VenueID:   7,
		DayID:     7,
		Salary:    170_000,
		StartTime: "13.00",
		EndTime:   "15.00",
	}

	venue31Days := entity.VenueDay{
		VenueID:   8,
		DayID:     7,
		Salary:    170_000,
		StartTime: "15.00",
		EndTime:   "17.00",
	}

	venue32Days := entity.VenueDay{
		VenueID:   8,
		DayID:     7,
		Salary:    170_000,
		StartTime: "19.00",
		EndTime:   "21.00",
	}

	venue33Days := entity.VenueDay{
		VenueID:   9,
		DayID:     1,
		Salary:    150_000,
		StartTime: "10.00",
		EndTime:   "12.00",
	}

	venue34Days := entity.VenueDay{
		VenueID:   9,
		DayID:     2,
		Salary:    150_000,
		StartTime: "13.00",
		EndTime:   "15.00",
	}

	venue35Days := entity.VenueDay{
		VenueID:   10,
		DayID:     1,
		Salary:    150_000,
		StartTime: "15.00",
		EndTime:   "17.00",
	}

	venue36Days := entity.VenueDay{
		VenueID:   10,
		DayID:     2,
		Salary:    150_000,
		StartTime: "19.00",
		EndTime:   "21.00",
	}

	venueDays = append(venueDays, venue1Days, venue2Days, venue3Days,
		venue4Days, venue5Days, venue6Days, venue7Days, venue8Days, venue9Days, venue10Days,
		venue11Days, venue12Days, venue13Days, venue14Days, venue15Days, venue16Days, venue17Days, venue18Days, venue19Days, venue20Days,
		venue21Days, venue22Days, venue23Days, venue24Days, venue25Days, venue26Days, venue27Days, venue28Days, venue29Days, venue30Days,
		venue31Days, venue32Days, venue33Days, venue34Days, venue35Days, venue36Days)

	for _, venueDay := range venueDays {
		if err := db.Create(&venueDay).Error; err != nil {
			return err
		}
	}
	return nil
}
