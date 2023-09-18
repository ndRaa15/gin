package repository

import (
	"context"

	"gin/internal/domain/entity"
	"gin/internal/infrastructure/mysql"
)

type InstrumentRepository struct {
	db *mysql.DB
}

func NewInstrumentRepository(db *mysql.DB) *InstrumentRepository {
	return &InstrumentRepository{db}
}

func (ir *InstrumentRepository) GetAllInstrument(ctx context.Context) ([]*entity.Instrument, error) {
	var instruments []*entity.Instrument
	if err := ir.db.Debug().WithContext(ctx).Find(&instruments).Error; err != nil {
		return nil, err
	}
	return instruments, nil
}

func (ir *InstrumentRepository) GetByID(ctx context.Context, id uint) (*entity.Instrument, error) {
	var instrument entity.Instrument
	if err := ir.db.Debug().WithContext(ctx).Where("id = ?", id).First(&instrument).Error; err != nil {
		return nil, err
	}
	return &instrument, nil
}

func (ir *InstrumentRepository) Update(ctx context.Context, instrument *entity.Instrument, id uint) (*entity.Instrument, error) {
	if err := ir.db.Debug().WithContext(ctx).Model(&instrument).Where("id = ?", id).Updates(&instrument).Error; err != nil {
		return nil, err
	}
	return instrument, nil
}

func (ir *InstrumentRepository) CreateRentInstrument(ctx context.Context, rentInstrument *entity.RentInstrument) (*entity.RentInstrument, error) {
	if err := ir.db.Debug().WithContext(ctx).Model(&rentInstrument).Create(&rentInstrument).Error; err != nil {
		return nil, err
	}
	return rentInstrument, nil
}

func SeedInstrument(db *mysql.DB) error {
	instrument1 := entity.Instrument{
		ID:            1,
		Name:          "Gitar Akustik",
		Owner:         "Chyntia Alex",
		ShortDesc:     "Gitar Cowboy Tipe GW-120",
		Description:   "Deskripsi Produk Gitar Akustik Merk Cowboy Original Tipe GW-120 NS NA Ukuran 3/4 Senar String Trusrod",
		RentPrice:     150_000,
		District:      "Klojen",
		City:          "Malang",
		Province:      "Jawa Timur",
		Street:        "Jalan Raya Klojen",
		Spesification: "Body: Sapele Wood (Brown) Spruce Wood , Neck: Sapele Wood , Fretboard: Mahogany , Bridge: Mahogany , Dryer: Steel Groover , Senar: Steel String 0.10 , Trusrod Dual Action , Dot Inlay,",
		OwnerNumber:   "+6282145278752",
		Weight:        1500,
		IsBooked:      false,
		Rating:        4.5,
		Photo:         "https://drive.google.com/uc?id=1Ggx8HYOrGXbqpxkl6sUkUOLk0xn4Ids5",
	}

	instrument2 := entity.Instrument{
		ID:            2,
		Name:          "Drum Akustik Rolling",
		Owner:         "Queen Studio",
		ShortDesc:     "Drum Akustik Junior Rolling JBJ1049A",
		Description:   "Drum Akustik Untuk Penggunaan Musik",
		RentPrice:     500_000,
		District:      "Singosari",
		City:          "Malang",
		Province:      "Jawa Timur",
		Street:        "Jalan Raya Asmaara",
		Spesification: "FLOOR: 12 , TOM 1: 8 , TOM2: 10 , SNARE: 10 , BASS DRUM: 16",
		OwnerNumber:   "+6282145278752",
		Weight:        5000,
		IsBooked:      false,
		Rating:        4.5,
		Photo:         "https://drive.google.com/uc?id=187tarfyKdHHX8Ttk7Fj4Px5QSAz2c8Ms",
	}

	instrument3 := entity.Instrument{
		ID:            3,
		Name:          "Gitar Elektrik Ibanez",
		Owner:         "Studio Musik Roy",
		ShortDesc:     "Ibanez GRX20-BKN",
		Description:   "GRX20 adalah sebuah gitar elektrik solid body seri RX yang diperkenalkan Ibanez pada tahun 2000. Diproduksi di China dan Indonesia sebagai bagian dari lini entry-level GIO. GRX20 memiliki dua humbucking pickups yang dipasang di pickguard dan sebuah tremolo bridge singkronisasi dua titik.",
		RentPrice:     150_000,
		District:      "Tlogomas",
		City:          "Malang",
		Province:      "Jawa Timur",
		Street:        "Jalan Raya Mendut",
		Spesification: "Body Material: Poplar , Fretboard Material: Treated New Zealand Pine , Fretboard Radius: 400mmR , Scale Length: 648mm/25.5 , Fretboard Inlays: WHite Dot , Pickup(s): Infinity R (H) (Passive/Ceramic) , Bridge: Ibanez T106 tremolo , Strings: .009/.011/.016/.024/.032/.042 , Hardware Finish: Chrome",
		OwnerNumber:   "+6282145278752",
		Weight:        1500,
		IsBooked:      false,
		Rating:        4.5,
		Photo:         "https://drive.google.com/uc?id=1r9G7zWlOxDHcJx7Kh94DBVHQyWwGEFiu",
	}

	instrument4 := entity.Instrument{
		ID:            4,
		Name:          "Cajoon Pearl",
		Owner:         "Billy Musik",
		ShortDesc:     "Cajon Pearl Primero PBC123B TR",
		Description:   "Kajoon dengan port bass yang terintegrasi dan rangkaian kawat snare yang tetap, Pearl Primero Cajon menghadirkan beragam suara perkusi yang lebih luas dibandingkan dengan banyak cajon lainnya. Suara bassnya dalam dan kaya. Dan kawat snare menambahkan kecerahan dan artikulasi, memberikan drum kotak Anda kilauan yang diperlukan untuk tetap berada dalam mix.",
		RentPrice:     100_000,
		District:      "Lowokwaru",
		City:          "Malang",
		Province:      "Jawa Timur",
		Street:        "Jalan Raya Singosari",
		Spesification: "Type: Cajon Snare , Size: 11.75 x 11.75 x 19.25 , Material: MDF with Meranti faceplate, Manufacturer Part Number: PBC123BTR",
		OwnerNumber:   "+6282145278752",
		Weight:        2000,
		IsBooked:      false,
		Rating:        4.3,
		Photo:         "https://drive.google.com/uc?id=1FRQyg-_FtfS96_0pFQMNKvE4MzU7DQ5Q",
	}

	instrument5 := entity.Instrument{
		ID:            5,
		Name:          "Trumpet Cowboy",
		Owner:         "Tom Studio ",
		ShortDesc:     "Trumpet Merk Cowboy Original Warna Gold",
		Description:   "Trompet adalah salah satu jenis alat musik tiup logam yang populer. Alat musik ini terdiri dari pipa logam melengkung dengan corong berbentuk corolla di salah satu ujungnya dan katup di sisi lainnya. Trompet biasanya dimainkan dengan meniup udara melalui corongnya sambil mengatur nada dan nada dengan menekan atau melepaskan katupnya.",
		RentPrice:     300_000,
		District:      "Lowokwaru",
		City:          "Malang",
		Province:      "Jawa Timur",
		Street:        "Jalan Raya Borobudur",
		Spesification: "Warna: Gold , Bahan: Gold Brass Leadpipe (Nickel Gold) , Panjang: 55 cm , Tinggi : 15 cm , Diameter : 12,5 cm",
		OwnerNumber:   "+6282145278752",
		Weight:        1000,
		IsBooked:      false,
		Rating:        4.6,
		Photo:         "https://drive.google.com/uc?id=1qfPBJLBXK65WfgU8OLDyzZ6nMb9aUgbR",
	}

	instrument6 := entity.Instrument{
		ID:            6,
		Name:          "61 Key Portable Keyboard",
		Owner:         "Suka Studio",
		ShortDesc:     "Casio CT-X700 61-Key Portable Keyboard",
		Description:   "Kita sambut, CT-X700. Harga yang terjangkau menjadikannya pilihan yang tepat untuk pemain casual atau pemula, dan suaranya yang memukau menjadikannya pilihan wajib bagi pemain keyboard dari tingkat mana pun yang membutuhkan instrumen portabel.",
		RentPrice:     300_000,
		District:      "Kepanjen",
		City:          "Malang",
		Province:      "Jawa Timur",
		Street:        "Jalan Raya Prambanan",
		Spesification: "No. of keys: 61 , standard size keys , Touch Response: Sensitivity 3 types - Off , Maximum Polyphony: 48 notes (24 for certain tones) , Built-in Tones: 600 , Functions: Layer - Split - Piano/Organ button , Reverb: 1 to 20 - Off , Chorus: 1 to 10, Tone , Beats per Measure: 0 to 9 , Tempo Range: 20 to 255 , Demo Song: 1",
		OwnerNumber:   "+6282145278752",
		Weight:        7000,
		IsBooked:      true,
		Rating:        4.9,
		Photo:         "https://drive.google.com/uc?id=19wc7W3r1GtyxGp-P7mMYJ3So7nzgRU6H",
	}

	instrument7 := entity.Instrument{
		ID:            8,
		Name:          "Drum Set",
		Owner:         "Toko Musik Harmoni",
		ShortDesc:     "Drum Set Standar",
		Description:   "Drum set lengkap dengan bass drum, snare drum, tom-tom, hi-hat, cymbals, dan pedal. Cocok untuk pemula dan pemain tingkat menengah.",
		RentPrice:     300_000,
		District:      "Klojen",
		City:          "Malang",
		Province:      "Jawa Timur",
		Street:        "Jalan Pahlawan No. 456",
		Spesification: "Termasuk bass drum, snare drum, tom-tom, hi-hat, cymbals, dan pedal",
		OwnerNumber:   "+6281234567890",
		Weight:        2500,
		IsBooked:      false,
		Rating:        4.2,
		Photo:         "https://drive.google.com/uc?id=1IZ73AQXjW2HrjRlA6CHkZmBwjDKuDepQ",
	}

	instruments := []entity.Instrument{instrument1, instrument2, instrument3, instrument4, instrument5, instrument6, instrument7}
	for _, instrument := range instruments {
		if err := db.Debug().Create(&instrument).Error; err != nil {
			return err
		}
	}
	return nil
}
