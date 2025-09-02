INSERT INTO gaji (jabatan, gaji_pokok) VALUES
('Junior', 5000000),
('Senior', 8000000),
('Manager', 15000000),
('Excecutif', 25000000),
('Chief', 35000000);

INSERT INTO karyawan (nama, email, nomor_hp) VALUES
('Budi Santoso', 'budi.s@example.com', '081234567890'),
('Dewi Lestari', 'dewi.l@example.com', '081345678901'),
('Agus Wijaya', 'agus.w@example.com', '081456789012'),
('Siti Rahayu', 'siti.r@example.com', '081567890123'),
('Joko Susilo', 'joko.s@example.com', '081678901234'),
('Ratna Permata', 'ratna.p@example.com', '081789012345'),
('Andi Prasetyo', 'andi.p@example.com', '081890123456'),
('Eka Putri', 'eka.p@example.com', '081901234567'),
('Bayu Nugroho', 'bayu.n@example.com', '081912345678'),
('Lina Kartika', 'lina.k@example.com', '081923456789'),
('Deni Pratama', 'deni.p@example.com', '081934567890'),
('Nina Adelia', 'nina.a@example.com', '081945678901'),
('Rudi Hermawan', 'rudi.h@example.com', '081956789012'),
('Fina Maharani', 'fina.m@example.com', '081967890123'),
('Gani Ramadhan', 'gani.r@example.com', '081978901234');

INSERT INTO kehadiran (id_kehadiran, id_karyawan, bulan, tahun, hadir, alpa, sakit, izin) VALUES
(1, 1, 'September', 2025, 20, 1, 1, 0),
(2, 2, 'September', 2025, 18, 2, 1, 1),
(3, 3, 'September', 2025, 22, 0, 0, 0);