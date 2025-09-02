CREATE TABLE karyawan (
    id_karyawan INT(11) NOT NULL AUTO_INCREMENT,
    nama VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    nomor_hp VARCHAR(20) NOT NULL,
    PRIMARY KEY (id_karyawan)
);

CREATE TABLE gaji_pokok (
    id_gaji INT(11) NOT NULL AUTO_INCREMENT,
    jabatan VARCHAR(100) NOT NULL,
    gaji_pokok INT(11) NOT NULL,
    PRIMARY KEY (id_gaji)
);

CREATE TABLE kehadiran (
    id_kehadiran INT(11) NOT NULL AUTO_INCREMENT,
    id_karyawan INT(11) NOT NULL,
    bulan VARCHAR(20) NOT NULL,
    tahun INT(4) NOT NULL,
    hadir INT(11) NOT NULL,
    alpa INT(11) NOT NULL,
    sakit INT(11) NOT NULL,
    izin INT(11) NOT NULL,
    PRIMARY KEY (id_kehadiran),
    FOREIGN KEY (id_karyawan) REFERENCES karyawan(id_karyawan) ON DELETE CASCADE
);

CREATE TABLE status_pembayaran_gaji (
    id_pembayaran INT(11) NOT NULL AUTO_INCREMENT,
    id_karyawan INT(11) NOT NULL,
    id_gaji INT(11) NOT NULL,
    tanggal_pembayaran DATE NOT NULL,
    status ENUM('Pending', 'Done') NOT NULL,
    status_prestasi VARCHAR(100),
    tunjangan_prestasi INT(11) DEFAULT 0,
    jenis_pelanggaran TEXT,
    potongan_gaji INT(11) DEFAULT 0,
    total_gaji INT(11) NOT NULL,
    PRIMARY KEY (id_pembayaran),
    FOREIGN KEY (id_karyawan) REFERENCES karyawan(id_karyawan) ON DELETE CASCADE,
    FOREIGN KEY (id_gaji) REFERENCES gaji_pokok(id_gaji) ON DELETE CASCADE
);