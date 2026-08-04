package models

import (
	"log"
	"time"

	"backend-ocean-basketball/pkg/utils"
	"gorm.io/gorm"
)

func SeedDatabase(db *gorm.DB) {
	log.Println("Starting comprehensive database seeding...")

	// Clean up any empty ID records
	db.Exec("DELETE FROM coaches WHERE id = '' OR id IS NULL")
	db.Exec("DELETE FROM courts WHERE id = '' OR id IS NULL")
	db.Exec("DELETE FROM classes WHERE id = '' OR id IS NULL")
	db.Exec("DELETE FROM students WHERE id = '' OR id IS NULL")
	db.Exec("DELETE FROM banners WHERE id = '' OR id IS NULL")
	db.Exec("DELETE FROM faqs WHERE id = '' OR id IS NULL")
	db.Exec("DELETE FROM reviews WHERE id = '' OR id IS NULL")
	db.Exec("DELETE FROM tuition_plans WHERE id = '' OR id IS NULL")
	db.Exec("DELETE FROM tournaments WHERE id = '' OR id IS NULL")

	// 1. Seed Admin & Coach Users
	adminHash, _ := utils.HashPassword("admin123")
	coachHash, _ := utils.HashPassword("coach123")

	users := []User{
		{
			ID:        "u-admin-1",
			Name:      "Quản trị viên Ocean",
			Email:     "admin@oceanbasketball.vn",
			Password:  adminHash,
			Role:      "admin",
			Avatar:    "/images/img1.jpg",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        "u-coach-1",
			Name:      "HLV Nguyễn Văn Hùng",
			Email:     "hung.nv@oceanbasketball.vn",
			Password:  coachHash,
			Role:      "coach",
			Avatar:    "/images/img2.jpg",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        "u-coach-2",
			Name:      "HLV Trần Đức Anh",
			Email:     "ducanh.tran@oceanbasketball.vn",
			Password:  coachHash,
			Role:      "coach",
			Avatar:    "/images/img3.jpg",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	for _, u := range users {
		var count int64
		db.Model(&User{}).Where("id = ? OR email = ?", u.ID, u.Email).Count(&count)
		if count == 0 {
			db.Create(&u)
		}
	}
	log.Println("✅ Seeded Users")

	// 2. Seed Coaches
	coaches := []Coach{
		{
			ID:             "c-1",
			Name:           "Nguyễn Văn Hùng",
			Email:          "hung.nv@oceanbasketball.vn",
			Phone:          "0912345678",
			Avatar:         "/images/img2.jpg",
			Specialization: "HLV Trưởng - Kỹ thuật cơ bản & Tư duy chiến thuật",
			Experience:     "8 năm kinh nghiệm huấn luyện trẻ em & U16",
			Achievements:   []string{"Vô địch Giải Trẻ Hà Nội 2023", "Chứng chỉ HLV FIBA Level 2", "Huy chương Bạc Giải Bóng rổ Học sinh"},
			Bio:            "Cựu vận động viên bóng rổ quốc gia, tâm huyết với sự phát triển kỹ năng và thể chất của thế hệ trẻ.",
			IsActive:       true,
		},
		{
			ID:             "c-2",
			Name:           "Trần Đức Anh",
			Email:          "ducanh.tran@oceanbasketball.vn",
			Phone:          "0987654321",
			Avatar:         "/images/img3.jpg",
			Specialization: "HLV Thể lực & Kỹ năng ném rổ dứt điểm",
			Experience:     "6 năm huấn luyện chuyên sâu",
			Achievements:   []string{"Huy chương Vàng VBA Junior 2024", "Chứng chỉ Thể lực Quốc tế CSCS"},
			Bio:            "Chuyên gia phát triển thể lực, sức bật và khả năng dứt điểm ném rổ chính xác cho học viên.",
			IsActive:       true,
		},
		{
			ID:             "c-3",
			Name:           "Lê Hoàng Nam",
			Email:          "nam.lh@oceanbasketball.vn",
			Phone:          "0933445566",
			Avatar:         "/images/img4.jpg",
			Specialization: "HLV Kiểm soát bóng & Phòng thủ đồng đội",
			Experience:     "5 năm kinh nghiệm",
			Achievements:   []string{"Chứng chỉ HLV Thể thao Quốc gia", "Giải Nhì VBA 3x3 2023"},
			Bio:            "Tập trung nâng cao tư duy đọc trận đấu, kỹ năng dắt bóng và phản xạ phòng thủ cá nhân.",
			IsActive:       true,
		},
		{
			ID:             "c-4",
			Name:           "Phạm Quốc Huy",
			Email:          "huy.pq@oceanbasketball.vn",
			Phone:          "0944556677",
			Avatar:         "/images/img14.jpg",
			Specialization: "HLV Phát triển Kỹ năng Cá nhân U8-U12",
			Experience:     "4 năm kinh nghiệm",
			Achievements:   []string{"HLV Xuất sắc Mùa giải 2025", "Vô địch U12 Hanoi Youth League"},
			Bio:            "Truyền cảm hứng và tinh thần đồng đội cho học viên mầm non & tiểu học.",
			IsActive:       true,
		},
		{
			ID:             "c-5",
			Name:           "Vũ Minh Đức",
			Email:          "duc.vm@oceanbasketball.vn",
			Phone:          "0955667788",
			Avatar:         "/images/img15.jpg",
			Specialization: "HLV Chuyên sâu 3x3 & Đột phá",
			Experience:     "7 năm thi đấu & giảng dạy",
			Achievements:   []string{"Chứng chỉ HLV FIBA 3x3", "Huy chương Vàng Giải Sinh viên Hà Nội"},
			Bio:            "Chuyên huấn luyện các bài tập dứt điểm cận rổ, đảo người và phản công nhanh.",
			IsActive:       true,
		},
	}
	for _, c := range coaches {
		var count int64
		db.Model(&Coach{}).Where("id = ? OR email = ?", c.ID, c.Email).Count(&count)
		if count == 0 {
			db.Create(&c)
		}
	}
	log.Println("✅ Seeded Coaches")

	// 3. Seed Courts
	courts := []Court{
		{
			ID:         "court-1",
			Name:       "Sân Bóng Rổ Vinhomes Ocean Park 1",
			Address:    "Khu Biệt thự Ngọc Trai, Vinhomes Ocean Park 1, Gia Lâm, Hà Nội",
			Latitude:   21.0035,
			Longitude:  105.9520,
			Images:     []string{"/images/img5.jpg", "/images/img6.jpg", "/images/img7.jpg", "/images/img8.jpg"},
			Facilities: []string{"Mặt sân thảm chuẩn FIBA", "Hệ thống đèn chiếu sáng cao cấp", "Phòng thay đồ khép kín", "Khán đài 500 chỗ", "Khu phụ huynh chờ có điều hòa"},
			ClassCount: 8,
		},
		{
			ID:         "court-2",
			Name:       "Sân Bóng Rổ Vinhomes Ocean Park 2",
			Address:    "Phân khu Chà Là, Vinhomes Ocean Park 2, Hưng Yên",
			Latitude:   20.9780,
			Longitude:  105.9810,
			Images:     []string{"/images/img9.jpg", "/images/img10.jpg", "/images/img11.jpg"},
			Facilities: []string{"Sân trong nhà mái che", "Bảng rổ Kính cường lực cao cấp", "Điều hòa trung tâm", "Phòng tắm nước nóng"},
			ClassCount: 5,
		},
		{
			ID:         "court-3",
			Name:       "Sân Bóng Rổ Vinhomes Ocean Park 3",
			Address:    "Phân khu Phố Biển, Vinhomes Ocean Park 3, Hưng Yên",
			Latitude:   20.9650,
			Longitude:  105.9920,
			Images:     []string{"/images/img12.jpg", "/images/img13.jpg", "/images/img14.jpg"},
			Facilities: []string{"Sân ngoài trời thoáng mát", "Hệ thống lưới che nắng", "Khu phụ huynh nghỉ ngơi", "Wifi miễn phí"},
			ClassCount: 4,
		},
		{
			ID:         "court-4",
			Name:       "Sân Bóng Rổ Vinhomes Times City",
			Address:    "458 Minh Khai, Phường Vĩnh Tuy, Hai Bà Trưng, Hà Nội",
			Latitude:   20.9950,
			Longitude:  105.8670,
			Images:     []string{"/images/img15.jpg", "/images/img16.jpg"},
			Facilities: []string{"Sân thảm cao cấp", "Đèn chiếu đêm LED 1000W", "Hệ thống âm thanh sôi động"},
			ClassCount: 6,
		},
	}
	for _, ct := range courts {
		var count int64
		db.Model(&Court{}).Where("id = ?", ct.ID).Count(&count)
		if count == 0 {
			db.Create(&ct)
		}
	}
	log.Println("✅ Seeded Courts")

	// 4. Seed Classes & Schedules
	classes := []Class{
		{
			ID:          "cls-1",
			Name:        "Lớp U8 - U10 (Cơ bản Mầm non & Tiểu học)",
			CourtID:     "court-1",
			CoachID:     "c-1",
			Level:       "Cơ bản",
			Image:       "/images/img8.jpg",
			MaxStudents: 15,
		},
		{
			ID:          "cls-2",
			Name:        "Lớp U11 - U14 (Phát triển Kỹ năng THCS)",
			CourtID:     "court-1",
			CoachID:     "c-2",
			Level:       "Nâng cao",
			Image:       "/images/img9.jpg",
			MaxStudents: 18,
		},
		{
			ID:          "cls-3",
			Name:        "Lớp U15 - U18 (Chuyên sâu & Thi đấu THPT)",
			CourtID:     "court-2",
			CoachID:     "c-3",
			Level:       "Chuyên sâu",
			Image:       "/images/img10.jpg",
			MaxStudents: 20,
		},
		{
			ID:          "cls-4",
			Name:        "Lớp U10 - U12 (Phát triển Năng khiếu)",
			CourtID:     "court-3",
			CoachID:     "c-4",
			Level:       "Nâng cao",
			Image:       "/images/img11.jpg",
			MaxStudents: 16,
		},
		{
			ID:          "cls-5",
			Name:        "Lớp Đội Tuyển Trẻ OceanBasketball U16",
			CourtID:     "court-4",
			CoachID:     "c-5",
			Level:       "Chuyên nghiệp",
			Image:       "/images/img12.jpg",
			MaxStudents: 12,
		},
	}
	for _, cl := range classes {
		var count int64
		db.Model(&Class{}).Where("id = ?", cl.ID).Count(&count)
		if count == 0 {
			db.Create(&cl)
		}
	}

	schedules := []ClassSchedule{
		{ID: "sch-1", ClassID: "cls-1", DayOfWeek: 1, StartTime: "17:30", EndTime: "19:00"},
		{ID: "sch-2", ClassID: "cls-1", DayOfWeek: 3, StartTime: "17:30", EndTime: "19:00"},
		{ID: "sch-3", ClassID: "cls-1", DayOfWeek: 5, StartTime: "17:30", EndTime: "19:00"},
		{ID: "sch-4", ClassID: "cls-2", DayOfWeek: 2, StartTime: "18:00", EndTime: "19:30"},
		{ID: "sch-5", ClassID: "cls-2", DayOfWeek: 4, StartTime: "18:00", EndTime: "19:30"},
		{ID: "sch-6", ClassID: "cls-2", DayOfWeek: 6, StartTime: "08:00", EndTime: "09:30"},
		{ID: "sch-7", ClassID: "cls-3", DayOfWeek: 6, StartTime: "16:00", EndTime: "18:00"},
		{ID: "sch-8", ClassID: "cls-3", DayOfWeek: 0, StartTime: "16:00", EndTime: "18:00"},
	}
	for _, sch := range schedules {
		var count int64
		db.Model(&ClassSchedule{}).Where("id = ?", sch.ID).Count(&count)
		if count == 0 {
			db.Create(&sch)
		}
	}
	log.Println("✅ Seeded Classes & Schedules")

	// 5. Seed Students
	students := []Student{
		{
			ID:          "st-1",
			Name:        "Phạm Minh Trí",
			Avatar:      "/images/img15.jpg",
			BirthYear:   2015,
			ParentName:  "Phạm Văn Thắng",
			ParentPhone: "0911223344",
			ClassID:     "cls-1",
			Status:      "active",
			JoinDate:    time.Now().AddDate(0, -3, 0),
			Notes:       "Nhiệt tình, có tố chất kiểm soát bóng tốt",
		},
		{
			ID:          "st-2",
			Name:        "Nguyễn Hoàng Anh",
			Avatar:      "/images/img16.jpg",
			BirthYear:   2013,
			ParentName:  "Nguyễn Đức Hoàng",
			ParentPhone: "0922334455",
			ClassID:     "cls-2",
			Status:      "active",
			JoinDate:    time.Now().AddDate(0, -5, 0),
			Notes:       "Tốc độ di chuyển nhanh, dứt điểm tốt",
		},
		{
			ID:          "st-3",
			Name:        "Trần Việt Cường",
			Avatar:      "/images/img17.jpg",
			BirthYear:   2010,
			ParentName:  "Trần Văn Hùng",
			ParentPhone: "0933445566",
			ClassID:     "cls-3",
			Status:      "active",
			JoinDate:    time.Now().AddDate(0, -8, 0),
			Notes:       "Đội trưởng đội tuyển U16 CLB",
		},
		{
			ID:          "st-4",
			Name:        "Lê Bảo Nam",
			Avatar:      "/images/img18.jpg",
			BirthYear:   2014,
			ParentName:  "Lê Văn Tuấn",
			ParentPhone: "0944556677",
			ClassID:     "cls-1",
			Status:      "active",
			JoinDate:    time.Now().AddDate(0, -2, 0),
			Notes:       "Chăm chỉ tập luyện, tiến bộ rõ rệt",
		},
		{
			ID:          "st-5",
			Name:        "Đỗ Đức Minh",
			Avatar:      "/images/img19.jpg",
			BirthYear:   2011,
			ParentName:  "Đỗ Quốc Bảo",
			ParentPhone: "0955667788",
			ClassID:     "cls-2",
			Status:      "active",
			JoinDate:    time.Now().AddDate(0, -4, 0),
			Notes:       "Chơi vị trí Hậu vệ dẫn dắt tốt",
		},
		{
			ID:          "st-6",
			Name:        "Nguyễn Khánh Linh",
			Avatar:      "/images/img20.jpg",
			BirthYear:   2012,
			ParentName:  "Nguyễn Văn Bình",
			ParentPhone: "0966778899",
			ClassID:     "cls-4",
			Status:      "active",
			JoinDate:    time.Now().AddDate(0, -1, 0),
			Notes:       "Ném trung bình chính xác",
		},
	}
	for _, st := range students {
		var count int64
		db.Model(&Student{}).Where("id = ?", st.ID).Count(&count)
		if count == 0 {
			db.Create(&st)
		}
	}
	log.Println("✅ Seeded Students")

	// 6. Seed Trial Registrations
	trials := []TrialRegistration{
		{
			ID:               "trial-1",
			ParentName:       "Nguyễn Thanh Phong",
			ParentPhone:      "0977889900",
			StudentName:      "Nguyễn Phong Tiến",
			StudentBirthYear: 2016,
			PreferredCourt:   "Sân Vinhomes Ocean Park 1",
			Notes:            "Muốn đăng ký tập ca chiều 17:30",
			Status:           "pending",
			CreatedAt:        time.Now().AddDate(0, 0, -2),
		},
		{
			ID:               "trial-2",
			ParentName:       "Đỗ Ngọc Anh",
			ParentPhone:      "0988990011",
			StudentName:      "Đỗ Hoàng Lâm",
			StudentBirthYear: 2013,
			PreferredCourt:   "Sân Vinhomes Ocean Park 2",
			Notes:            "Bé đã từng tập bóng rổ 3 tháng ở trường",
			Status:           "approved",
			CreatedAt:        time.Now().AddDate(0, 0, -5),
		},
	}
	for _, tr := range trials {
		var count int64
		db.Model(&TrialRegistration{}).Where("id = ?", tr.ID).Count(&count)
		if count == 0 {
			db.Create(&tr)
		}
	}
	log.Println("✅ Seeded Trial Registrations")

	// 7. Seed Attendance Records
	attendances := []AttendanceRecord{
		{ID: "att-1", StudentID: "st-1", ClassID: "cls-1", Date: time.Now().AddDate(0, 0, -2), Status: "present", MarkedBy: "u-coach-1"},
		{ID: "att-2", StudentID: "st-4", ClassID: "cls-1", Date: time.Now().AddDate(0, 0, -2), Status: "present", MarkedBy: "u-coach-1"},
		{ID: "att-3", StudentID: "st-2", ClassID: "cls-2", Date: time.Now().AddDate(0, 0, -1), Status: "present", MarkedBy: "u-coach-2"},
		{ID: "att-4", StudentID: "st-5", ClassID: "cls-2", Date: time.Now().AddDate(0, 0, -1), Status: "excused", MarkedBy: "u-coach-2"},
		{ID: "att-5", StudentID: "st-3", ClassID: "cls-3", Date: time.Now().AddDate(0, 0, -3), Status: "present", MarkedBy: "u-coach-2"},
	}
	for _, att := range attendances {
		var count int64
		db.Model(&AttendanceRecord{}).Where("id = ?", att.ID).Count(&count)
		if count == 0 {
			db.Create(&att)
		}
	}
	log.Println("✅ Seeded Attendance Records")

	// 8. Seed Banners
	banners := []Banner{
		{
			ID:        "b-1",
			Image:     "/images/img1.jpg",
			Title:     "KHƠI GỢI ĐAM MÊ BÓNG RỔ TRẺ EM",
			Subtitle:  "Môi trường đào tạo chuẩn quốc tế tại Vinhomes Ocean Park với đội ngũ HLV giàu kinh nghiệm.",
			CTAText:   "Đăng ký học thử ngay",
			CTALink:   "/hoc-thu",
			Order:     1,
			IsActive:  true,
		},
		{
			ID:        "b-2",
			Image:     "/images/img2.jpg",
			Title:     "PHÁT TRIỂN THỂ LỰC & TƯ DUY ĐỒNG ĐỘI",
			Subtitle:  "Chương trình huấn luyện bài bản từ cơ bản đến nâng cao cho lứa tuổi 5 - 18 tuổi.",
			CTAText:   "Khám phá lộ trình",
			CTALink:   "/lo-trinh-tap-luyen",
			Order:     2,
			IsActive:  true,
		},
		{
			ID:        "b-3",
			Image:     "/images/img3.jpg",
			Title:     "GIẢI ĐẤU OCEAN BASKETBALL LEAGUE 2026",
			Subtitle:  "Cơ hội cọ xát, thi đấu đỉnh cao và thể hiện tài năng dành cho các học viên xuất sắc.",
			CTAText:   "Xem thông tin giải đấu",
			CTALink:   "/thi-dau",
			Order:     3,
			IsActive:  true,
		},
	}
	for _, bn := range banners {
		var count int64
		db.Model(&Banner{}).Where("id = ?", bn.ID).Count(&count)
		if count == 0 {
			db.Create(&bn)
		}
	}
	log.Println("✅ Seeded Banners")

	// 9. Seed Tuition Plans
	plans := []TuitionPlan{
		{
			ID:              "tp-1",
			Name:            "Gói Trải Nghiệm (1 Tháng)",
			Price:           1200000,
			Duration:        "1 Tháng",
			SessionsPerWeek: 2,
			Features:        []string{"8 buổi tập trực tiếp cùng HLV", "Tặng 01 bộ đồng phục thi đấu", "Miễn phí 01 buổi học thử trải nghiệm"},
			IsPopular:       false,
		},
		{
			ID:              "tp-2",
			Name:            "Gói Tiêu Chuẩn (3 Tháng)",
			Price:           3200000,
			Duration:        "3 Tháng",
			SessionsPerWeek: 2,
			Features:        []string{"24 buổi tập kỹ thuật chuyên sâu", "Tặng 01 bộ đồng phục + 01 quả bóng rổ FIBA", "Giảm 10% khi đăng ký theo nhóm 2 người", "Tham gia thi đấu nội bộ hàng tháng"},
			IsPopular:       true,
		},
		{
			ID:              "tp-3",
			Name:            "Gói Chuyên Nghiệp (6 Tháng)",
			Price:           5900000,
			Duration:        "6 Tháng",
			SessionsPerWeek: 3,
			Features:        []string{"72 buổi tập nâng cao & thể lực", "Tặng full combo đồng phục + balo + bóng rổ", "Đánh giá chỉ số thể chất định kỳ 2 tháng/lần", "Ưu tiên tuyển chọn vào đội tuyển tham gia giải đấu"},
			IsPopular:       false,
		},
		{
			ID:              "tp-4",
			Name:            "Gói VIP Đào Tạo Vận Động Viên (1 Năm)",
			Price:           10800000,
			Duration:        "12 Tháng",
			SessionsPerWeek: 3,
			Features:        []string{"144 buổi tập chuyên sâu đỉnh cao", "Huấn luyện 1:1 cá nhân hóa với HLV Trưởng", "Tặng full trang thiết bị & giày thi đấu", "Đảm bảo suất thi đấu giải chính thức Hà Nội"},
			IsPopular:       false,
		},
	}
	for _, tp := range plans {
		var count int64
		db.Model(&TuitionPlan{}).Where("id = ?", tp.ID).Count(&count)
		if count == 0 {
			db.Create(&tp)
		}
	}
	log.Println("✅ Seeded Tuition Plans")

	// 10. Seed FAQs
	faqs := []FAQ{
		{
			ID:       "faq-1",
			Question: "Độ tuổi nào có thể tham gia các lớp học tại OceanBasketball?",
			Answer:   "OceanBasketball tiếp nhận học viên từ 5 đến 18 tuổi, được phân chia thành các nhóm tuổi U8, U10, U14, U18 để đảm bảo giáo trình phù hợp nhất với thể trạng.",
			Category: "Chung",
			Order:    1,
		},
		{
			ID:       "faq-2",
			Question: "Học viên chưa từng chơi bóng rổ có theo học được không?",
			Answer:   "Hoàn toàn được! Các lớp U8 và lớp Cơ bản được thiết kế riêng dành cho các bạn mới bắt đầu để rèn luyện tư thế, kỹ thuật nhồi bóng và ném rổ chuẩn.",
			Category: "Khóa học",
			Order:    2,
		},
		{
			ID:       "faq-3",
			Question: "Trung tâm có hỗ trợ học thử miễn phí không?",
			Answer:   "OceanBasketball hỗ trợ 01 buổi học thử miễn phí 100% để phụ huynh và các em trải nghiệm không khí tập luyện trước khi chính thức đăng ký.",
			Category: "Đăng ký",
			Order:    3,
		},
		{
			ID:       "faq-4",
			Question: "Nếu học viên nghỉ học có được học bù không?",
			Answer:   "CLB hỗ trợ học bù vào các ca học khác cùng cấp độ trong vòng 30 ngày nếu phụ huynh xin nghỉ trước ca học ít nhất 2 tiếng.",
			Category: "Chính sách",
			Order:    4,
		},
	}
	for _, f := range faqs {
		var count int64
		db.Model(&FAQ{}).Where("id = ?", f.ID).Count(&count)
		if count == 0 {
			db.Create(&f)
		}
	}
	log.Println("✅ Seeded FAQs")

	// 11. Seed Reviews
	reviews := []Review{
		{
			ID:          "rev-1",
			ParentName:  "Chị Nguyễn Thu Trang",
			Avatar:      "/images/img18.jpg",
			Rating:      5,
			Content:     "Bé nhà mình học ở Ocean Park 1 được 6 tháng, con tự tin hẳn lên và chiều cao phát triển rất tốt. HLV Hùng rất kiên nhẫn và tận tâm!",
			StudentName: "Bé Minh Trí (8 tuổi)",
			IsVisible:   true,
			CreatedAt:   time.Now().AddDate(0, -1, 0),
		},
		{
			ID:          "rev-2",
			ParentName:  "Anh Hoàng Văn Nam",
			Avatar:      "/images/img19.jpg",
			Rating:      5,
			Content:     "Sân bãi sạch đẹp, an toàn. Thích nhất là trung tâm có các giải đấu nội bộ cho các con cọ xát hàng tháng.",
			StudentName: "Bé Hoàng Anh (11 tuổi)",
			IsVisible:   true,
			CreatedAt:   time.Now().AddDate(0, -2, 0),
		},
		{
			ID:          "rev-3",
			ParentName:  "Chị Trần Thanh Hà",
			Avatar:      "/images/img20.jpg",
			Rating:      5,
			Content:     "Giáo trình bài bản, môi trường năng động lành mạnh giúp con tránh xa điện thoại sau giờ học.",
			StudentName: "Bé Việt Cường (14 tuổi)",
			IsVisible:   true,
			CreatedAt:   time.Now().AddDate(0, -3, 0),
		},
	}
	for _, r := range reviews {
		var count int64
		db.Model(&Review{}).Where("id = ?", r.ID).Count(&count)
		if count == 0 {
			db.Create(&r)
		}
	}
	log.Println("✅ Seeded Reviews")

	// 12. Seed Tournaments
	tDate := time.Now().AddDate(0, 1, 0)
	tEndDate := tDate.AddDate(0, 0, 3)
	tournaments := []Tournament{
		{
			ID:          "t-1",
			Name:        "Ocean Basketball Summer Cup 2026",
			Description: "Giải đấu mùa hè tranh cúp vô địch giữa các phân khu Vinhomes Ocean Park 1, 2, 3.",
			Date:        tDate,
			EndDate:     &tEndDate,
			Location:    "Sân Bóng Rổ Vinhomes Ocean Park 1",
			Banner:      "/images/img20.jpg",
			Status:      "upcoming",
		},
		{
			ID:          "t-2",
			Name:        "Giải Bóng Rổ Trẻ Hà Nội Open 2026",
			Description: "Giải đấu quy tụ hơn 32 đội bóng học sinh trên toàn thành phố Hà Nội.",
			Date:        tDate.AddDate(0, 2, 0),
			Location:    "Sân Bóng Rổ Vinhomes Ocean Park 2",
			Banner:      "/images/img1.jpg",
			Status:      "upcoming",
		},
	}
	for _, t := range tournaments {
		var count int64
		db.Model(&Tournament{}).Where("id = ?", t.ID).Count(&count)
		if count == 0 {
			db.Create(&t)
		}
	}
	log.Println("✅ Seeded Tournaments")

	// 13. Seed Invoices & Payments
	dueDate := time.Now().AddDate(0, 0, 10)
	invoices := []Invoice{
		{
			ID:        "inv-1",
			StudentID: "st-1",
			Amount:    1200000,
			Month:     8,
			Year:      2026,
			Status:    "paid",
			DueDate:   dueDate,
			CreatedAt: time.Now().AddDate(0, 0, -5),
		},
		{
			ID:        "inv-2",
			StudentID: "st-2",
			Amount:    3200000,
			Month:     8,
			Year:      2026,
			Status:    "paid",
			DueDate:   dueDate,
			CreatedAt: time.Now().AddDate(0, 0, -5),
		},
		{
			ID:        "inv-3",
			StudentID: "st-3",
			Amount:    5900000,
			Month:     8,
			Year:      2026,
			Status:    "unpaid",
			DueDate:   dueDate,
			CreatedAt: time.Now().AddDate(0, 0, -2),
		},
	}
	for _, inv := range invoices {
		var count int64
		db.Model(&Invoice{}).Where("id = ?", inv.ID).Count(&count)
		if count == 0 {
			db.Create(&inv)
		}
	}

	payments := []Payment{
		{
			ID:        "pay-1",
			InvoiceID: "inv-1",
			Amount:    1200000,
			Method:    "transfer",
			Note:      "Thanh toán học phí tháng 8/2026 cho bé Minh Trí",
			CreatedAt: time.Now().AddDate(0, 0, -4),
		},
		{
			ID:        "pay-2",
			InvoiceID: "inv-2",
			Amount:    3200000,
			Method:    "cash",
			Note:      "Đã nộp tiền mặt tại văn phòng sân OceanPark 1",
			CreatedAt: time.Now().AddDate(0, 0, -3),
		},
	}
	for _, p := range payments {
		var count int64
		db.Model(&Payment{}).Where("id = ?", p.ID).Count(&count)
		if count == 0 {
			db.Create(&p)
		}
	}
	log.Println("✅ Seeded Invoices & Payments")

	// 14. Seed Contacts
	contacts := []Contact{
		{
			ID:        "ct-1",
			Name:      "Bùi Anh Tuấn",
			Email:     "tuan.ba@gmail.com",
			Phone:     "0909112233",
			Subject:   "Tư vấn lịch học U12 tại Ocean Park 2",
			Message:   "Cho mình hỏi lịch tập lớp U12 ở sân Chà Là chiều T7 & CN có còn trống chỗ không?",
			IsRead:    true,
			CreatedAt: time.Now().AddDate(0, 0, -3),
		},
		{
			ID:        "ct-2",
			Name:      "Đặng Phương Thảo",
			Email:     "thao.dp@outlook.com",
			Phone:     "0918223344",
			Subject:   "Đăng ký nhóm 3 học viên",
			Message:   "Mình muốn đăng ký cho 3 cháu cùng lứa tuổi U10, có được chiết khấu gói 3 tháng không?",
			IsRead:    false,
			CreatedAt: time.Now().AddDate(0, 0, -1),
		},
	}
	for _, ct := range contacts {
		var count int64
		db.Model(&Contact{}).Where("id = ?", ct.ID).Count(&count)
		if count == 0 {
			db.Create(&ct)
		}
	}
	log.Println("✅ Seeded Contacts")

	log.Println("🎉 ALL 16 TABLES SEEDED 100% SUCCESSFULLY!")
}
