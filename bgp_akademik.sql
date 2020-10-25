-- phpMyAdmin SQL Dump
-- version 4.9.5deb2
-- https://www.phpmyadmin.net/
--
-- Host: localhost:3306
-- Generation Time: Oct 25, 2020 at 09:00 PM
-- Server version: 10.3.24-MariaDB-2
-- PHP Version: 7.4.9

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `bgp_akademik`
--
CREATE DATABASE IF NOT EXISTS `bgp_akademik` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
USE `bgp_akademik`;

-- --------------------------------------------------------

--
-- Table structure for table `dosen`
--

CREATE TABLE `dosen` (
  `nip` varchar(100) NOT NULL,
  `nama` varchar(100) DEFAULT NULL,
  `jalan` varchar(100) DEFAULT NULL,
  `kecamatan` varchar(100) DEFAULT NULL,
  `kabupaten` varchar(100) DEFAULT NULL,
  `provinsi` varchar(100) DEFAULT NULL,
  `pangkat` varchar(100) DEFAULT NULL,
  `golongan` varchar(100) DEFAULT NULL,
  `email` varchar(100) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `dosen`
--

INSERT INTO `dosen` (`nip`, `nama`, `jalan`, `kecamatan`, `kabupaten`, `provinsi`, `pangkat`, `golongan`, `email`) VALUES
('111222333', 'Yunus Wijaya', 'Kenangan', 'Unknow', 'Padang', 'Sumbar', 'Pembina', 'IV.a', 'yunuswijaya@pnp.ac.id'),
('197804152000121002', 'Rahmat Hidayat, ST.,M.Sc.IT', 'Komp. Griya Tui Indah No. C-7, Belimbing', 'Kuranji', 'Padang', 'Sumbar', 'Penata Tingkat I', 'III.d', 'rahmathidayat@pnp.ac.id'),
('198104152006041002', 'Deddy Prayama, S.Kom.,M.ISD', 'Jl. Koto Panjang RT15/007', 'Pauh', 'Padang', 'Sumbar', 'Penata', 'III.c', 'deddy@pnp.ac.id'),
('198808252015041002', 'Alde Alanda, S.Kom.,MT', 'Komp. Unand Blok B/III/04/01 Gadut', 'Pauh', 'Padang', 'Sumbar', 'Penata Muda Tingkat I', 'III.b', 'alde_pnp@pnp.ac.id');

-- --------------------------------------------------------

--
-- Table structure for table `mahasiswa`
--

CREATE TABLE `mahasiswa` (
  `nobp` varchar(100) NOT NULL,
  `nama` varchar(100) DEFAULT NULL,
  `jalan` varchar(100) DEFAULT NULL,
  `kecamatan` varchar(100) DEFAULT NULL,
  `kabupaten` varchar(100) DEFAULT NULL,
  `provinsi` varchar(100) DEFAULT NULL,
  `jurusan` varchar(100) DEFAULT NULL,
  `prodi` varchar(100) DEFAULT NULL,
  `email` varchar(100) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `mahasiswa`
--

INSERT INTO `mahasiswa` (`nobp`, `nama`, `jalan`, `kecamatan`, `kabupaten`, `provinsi`, `jurusan`, `prodi`, `email`) VALUES
('1811081018', 'Salsabila Delaisy Permana', 'cempaka', 'tiakar', 'payakumbuh', 'sumbar', NULL, NULL, NULL),
('1811081035', 'Rozliyana ', 'rindu', 'yang terlupa', 'harapan', 'menunggu', NULL, NULL, NULL),
('1811082005', 'Ahmad Arif', 'Tan Malaka', 'Belubus', '50 Kota', 'Sumbar', 'TI', 'TRPL', NULL),
('1811082018', 'Bayu Gusti Paraya', 'Tan Malaka Km.12', 'Guguak', '50 Kota', 'Sumbar', 'Teknologi Informasi', 'D4-TRPL', 'bayu.gustip7820@gmail.com');

-- --------------------------------------------------------

--
-- Table structure for table `matkul`
--

CREATE TABLE `matkul` (
  `id` varchar(20) NOT NULL,
  `nama_mk` varchar(100) DEFAULT NULL,
  `sks` int(5) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `matkul`
--

INSERT INTO `matkul` (`id`, `nama_mk`, `sks`) VALUES
('MK008', 'Keamanan Jarkom', 2),
('MK011', 'Project 2 (Web)', 3),
('MK021', 'Web Framework', 3),
('MK022', 'Metode Penelitian', 2);

-- --------------------------------------------------------

--
-- Table structure for table `nilai`
--

CREATE TABLE `nilai` (
  `bp_mhs` varchar(100) NOT NULL,
  `id_matkul` varchar(20) NOT NULL,
  `nip_dosen` varchar(100) NOT NULL,
  `nilai` float NOT NULL,
  `semester` int(5) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `nilai`
--

INSERT INTO `nilai` (`bp_mhs`, `id_matkul`, `nip_dosen`, `nilai`, `semester`) VALUES
('1811082018', 'MK008', '198808252015041002', 3.5, 5),
('1811082018', 'MK011', '111222333', 3.75, 5),
('1811082018', 'MK021', '198104152006041002', 3.75, 5),
('1811082018', 'MK022', '197804152000121002', 4, 5),
('1811081018', 'MK008', '198808252015041002', 3.5, 5),
('1811081035', 'MK011', '111222333', 3, 5),
('1811081035', 'MK022', '197804152000121002', 4, 5),
('1811081018', 'MK021', '198104152006041002', 4, 5),
('1811082005', 'MK011', '111222333', 4, 5),
('1811082005', 'MK021', '198104152006041002', 3.5, 5),
('1811082005', 'MK022', '197804152000121002', 3.5, 5);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `dosen`
--
ALTER TABLE `dosen`
  ADD PRIMARY KEY (`nip`);

--
-- Indexes for table `mahasiswa`
--
ALTER TABLE `mahasiswa`
  ADD PRIMARY KEY (`nobp`);

--
-- Indexes for table `matkul`
--
ALTER TABLE `matkul`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `nilai`
--
ALTER TABLE `nilai`
  ADD KEY `bp_mhs` (`bp_mhs`,`id_matkul`,`nip_dosen`),
  ADD KEY `nip_dosen` (`nip_dosen`),
  ADD KEY `id_matkul` (`id_matkul`);

--
-- Constraints for dumped tables
--

--
-- Constraints for table `nilai`
--
ALTER TABLE `nilai`
  ADD CONSTRAINT `nilai_ibfk_1` FOREIGN KEY (`nip_dosen`) REFERENCES `dosen` (`nip`),
  ADD CONSTRAINT `nilai_ibfk_2` FOREIGN KEY (`id_matkul`) REFERENCES `matkul` (`id`),
  ADD CONSTRAINT `nilai_ibfk_3` FOREIGN KEY (`bp_mhs`) REFERENCES `mahasiswa` (`nobp`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
