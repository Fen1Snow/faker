package locales

var Vi = map[string]interface{}{
	"lorem": map[string]interface{}{
		"words": []string{
			"đã", "đang", "ừ", "ờ", "á", "không", "biết", "gì", "hết", "đâu", "nha", "thế", "thì", "là", "đánh", "đá", "đập", "phá", "viết", "vẽ", "tô", "thuê", "mướn", "mượn", "mua", "một", "hai", "ba", "bốn", "năm", "sáu", "bảy", "tám", "chín", "mười", "thôi", "việc", "nghỉ", "làm", "nhà", "cửa", "xe", "đạp", "ác", "độc", "khoảng", "khoan", "thuyền", "tàu", "bè", "lầu", "xanh", "đỏ", "tím", "vàng", "kim", "chỉ", "khâu", "may", "vá", "em", "anh", "yêu", "thương", "thích", "con", "cái", "bàn", "ghế", "tủ", "quần", "áo", "nón", "dép", "giày", "lỗi", "được", "ghét", "giết", "chết", "hết", "tôi", "bạn", "tui", "trời", "trăng", "mây", "gió", "máy", "hàng", "hóa", "leo", "núi", "bơi", "biển", "chìm", "xuồng", "nước", "ngọt", "ruộng", "đồng", "quê", "hương",
		},
	},
	"address": map[string]interface{}{
		"city_root": []string{
			"Bắc Giang", "Bắc Kạn", "Bắc Ninh", "Cao Bằng", "Điện Biên", "Hà Giang", "Hà Nam", "Hà Tây", "Hải Dương", "TP Hải Phòng", "Hòa Bình", "Hưng Yên", "Lai Châu", "Lào Cai", "Lạng Sơn", "Nam Định", "Ninh Bình", "Phú Thọ", "Quảng Ninh", "Sơn La", "Thái Bình", "Thái Nguyên", "Tuyên Quang", "Vĩnh Phúc", "Yên Bái", "TP Đà Nẵng", "Bình Định", "Đắk Lắk", "Đắk Nông", "Gia Lai", "Hà Tĩnh", "Khánh Hòa", "Kon Tum", "Nghệ An", "Phú Yên", "Quảng Bình", "Quảng Nam", "Quảng Ngãi", "Quảng Trị", "Thanh Hóa", "Thừa Thiên Huế", "TP TP. Hồ Chí Minh", "An Giang", "Bà Rịa Vũng Tàu", "Bạc Liêu", "Bến Tre", "Bình Dương", "Bình Phước", "Bình Thuận", "Cà Mau", "TP Cần Thơ", "Đồng Nai", "Đồng Tháp", "Hậu Giang", "Kiên Giang", "Lâm Đồng", "Long An", "Ninh Thuận", "Sóc Trăng", "Tây Ninh", "Tiền Giang", "Trà Vinh", "Vĩnh Long",
		},
		"city": []string{
			"#{city_root}",
		},
		"postcode": "/[A-PR-UWYZ0-9][A-HK-Y0-9][AEHMNPRTVXY0-9]?[ABEHMNPRVWXY0-9]? {1,2}[0-9][ABD-HJLN-UW-Z]{2}/",
		"county": []string{
			"Avon", "Bedfordshire", "Berkshire", "Borders", "Buckinghamshire", "Cambridgeshire", "Central", "Cheshire", "Cleveland", "Clwyd", "Cornwall", "County Antrim", "County Armagh", "County Down", "County Fermanagh", "County Londonderry", "County Tyrone", "Cumbria", "Derbyshire", "Devon", "Dorset", "Dumfries and Galloway", "Durham", "Dyfed", "East Sussex", "Essex", "Fife", "Gloucestershire", "Grampian", "Greater Manchester", "Gwent", "Gwynedd County", "Hampshire", "Herefordshire", "Hertfordshire", "Highlands and Islands", "Humberside", "Isle of Wight", "Kent", "Lancashire", "Leicestershire", "Lincolnshire", "Lothian", "Merseyside", "Mid Glamorgan", "Norfolk", "North Yorkshire", "Northamptonshire", "Northumberland", "Nottinghamshire", "Oxfordshire", "Powys", "Rutland", "Shropshire", "Somerset", "South Glamorgan", "South Yorkshire", "Staffordshire", "Strathclyde", "Suffolk", "Surrey", "Tayside", "Tyne and Wear", "Việt Nam", "Warwickshire", "West Glamorgan", "West Midlands", "West Sussex", "West Yorkshire", "Wiltshire", "Worcestershire",
		},
		"default_country": []string{
			"Việt Nam",
		},
	},
	"internet": map[string]interface{}{
		"domain_suffix": []string{
			"com", "net", "info", "vn", "com.vn",
		},
	},
	"phone_number": map[string]interface{}{
		"formats": []string{
			"01#### #####", "01### ######", "01#1 ### ####", "011# ### ####", "02# #### ####", "03## ### ####", "055 #### ####", "056 #### ####", "0800 ### ####", "08## ### ####", "09## ### ####", "016977 ####", "01### #####", "0500 ######", "0800 ######",
		},
	},
	"cell_phone": map[string]interface{}{
		"formats": []string{
			"074## ######", "075## ######", "076## ######", "077## ######", "078## ######", "079## ######",
		},
	},
	"name": map[string]interface{}{
		"first_name": []string{
			"Phạm", "Nguyễn", "Trần", "Lê", "Lý", "Hoàng", "Phan", "Vũ", "Tăng", "Đặng", "Bùi", "Đỗ", "Hồ", "Ngô", "Dương", "Đào", "Đoàn", "Vương", "Trịnh", "Đinh", "Lâm", "Phùng", "Mai", "Tô", "Trương", "Hà",
		},
		"last_name": []string{
			"Nam", "Trung", "Thanh", "Thị", "Văn", "Dương", "Tăng", "Quốc", "Như", "Phạm", "Nguyễn", "Trần", "Lê", "Lý", "Hoàng", "Phan", "Vũ", "Tăng", "Đặng", "Bùi", "Đỗ", "Hồ", "Ngô", "Dương", "Đào", "Đoàn", "Vương", "Trịnh", "Đinh", "Lâm", "Phùng", "Mai", "Tô", "Trương", "Hà", "Vinh", "Nhung", "Hòa", "Tiến", "Tâm", "Bửu", "Loan", "Hiền", "Hải", "Vân", "Kha", "Minh", "Nhân", "Triệu", "Tuân", "Hữu", "Đức", "Phú", "Khoa", "Thắgn", "Sơn", "Dung", "Tú", "Trinh", "Thảo", "Sa", "Kim", "Long", "Thi", "Cường", "Ngọc", "Sinh", "Khang", "Phong", "Thắm", "Thu", "Thủy", "Nhàn",
		},
		"name": []string{
			"#{first_name} #{last_name}", "#{first_name} #{last_name} #{last_name}", "#{first_name} #{last_name} #{last_name} #{last_name}",
		},
	},
	"company": map[string]interface{}{
		"prefix": []string{
			"Công ty", "Cty TNHH", "Cty", "Cửa hàng", "Trung tâm", "Chi nhánh",
		},
		"name": []string{
			"#{prefix} #{Name.last_name}",
		},
	},
}
