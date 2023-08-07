-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Aug 07, 2023 at 02:29 PM
-- Server version: 10.4.28-MariaDB
-- PHP Version: 8.0.28

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `mcdm`
--

-- --------------------------------------------------------

--
-- Table structure for table `alternatifs`
--

CREATE TABLE `alternatifs` (
  `id` bigint(20) NOT NULL,
  `kode` varchar(255) DEFAULT NULL,
  `kode_kriteria` varchar(255) DEFAULT NULL,
  `nama_alternatif` varchar(255) DEFAULT NULL,
  `nilai` varchar(255) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `alternatifs`
--

INSERT INTO `alternatifs` (`id`, `kode`, `kode_kriteria`, `nama_alternatif`, `nilai`, `created_at`) VALUES
(52, 'A1', 'C1', 'Pantai Panndan', '3', '2023-07-26 13:40:22.466'),
(53, 'A1', 'C4', 'Pantai Panndan', '3', '2023-07-26 13:40:22.467'),
(54, 'A1', 'C2', 'Pantai Panndan', '2', '2023-07-26 13:40:22.466'),
(55, 'A1', 'C6', 'Pantai Panndan', '4', '2023-07-26 13:40:22.469'),
(56, 'A1', 'C3', 'Pantai Panndan', '3', '2023-07-26 13:40:22.467'),
(57, 'A1', 'C5', 'Pantai Panndan', '3', '2023-07-26 13:40:22.469'),
(58, 'A1', 'C7', 'Pantai Panndan', '3', '2023-07-26 13:40:22.650'),
(59, 'A1', 'C8', 'Pantai Panndan', '4', '2023-07-26 13:40:22.676'),
(60, 'A2', 'C1', 'Muara Tapanuli Utara', '4', '2023-07-26 13:42:09.171'),
(61, 'A2', 'C2', 'Muara Tapanuli Utara', '4', '2023-07-26 13:42:09.171'),
(62, 'A2', 'C4', 'Muara Tapanuli Utara', '3', '2023-07-26 13:42:09.174'),
(63, 'A2', 'C3', 'Muara Tapanuli Utara', '2', '2023-07-26 13:42:09.172'),
(64, 'A2', 'C5', 'Muara Tapanuli Utara', '4', '2023-07-26 13:42:09.174'),
(65, 'A2', 'C6', 'Muara Tapanuli Utara', '2', '2023-07-26 13:42:09.174'),
(66, 'A2', 'C7', 'Muara Tapanuli Utara', '3', '2023-07-26 13:42:09.578'),
(67, 'A2', 'C8', 'Muara Tapanuli Utara', '3', '2023-07-26 13:42:09.578'),
(68, 'A3', 'C1', 'Kebun Raya Tongkoh', '3', '2023-07-26 13:42:55.783'),
(69, 'A3', 'C2', 'Kebun Raya Tongkoh', '3', '2023-07-26 13:42:55.784'),
(70, 'A3', 'C3', 'Kebun Raya Tongkoh', '4', '2023-07-26 13:42:55.784'),
(71, 'A3', 'C6', 'Kebun Raya Tongkoh', '4', '2023-07-26 13:42:55.788'),
(72, 'A3', 'C5', 'Kebun Raya Tongkoh', '2', '2023-07-26 13:42:55.788'),
(73, 'A3', 'C4', 'Kebun Raya Tongkoh', '2', '2023-07-26 13:42:55.787'),
(74, 'A3', 'C7', 'Kebun Raya Tongkoh', '2', '2023-07-26 13:42:56.157'),
(75, 'A3', 'C8', 'Kebun Raya Tongkoh', '2', '2023-07-26 13:42:56.201'),
(76, 'A4', 'C1', 'Sipinsur', '4', '2023-07-26 15:02:55.658'),
(77, 'A4', 'C2', 'Sipinsur', '2', '2023-07-26 15:02:55.658'),
(78, 'A4', 'C3', 'Sipinsur', '4', '2023-07-26 15:02:55.658'),
(79, 'A4', 'C4', 'Sipinsur', '3', '2023-07-26 15:02:55.659'),
(80, 'A4', 'C5', 'Sipinsur', '4', '2023-07-26 15:02:55.661'),
(81, 'A4', 'C6', 'Sipinsur', '2', '2023-07-26 15:02:55.661'),
(82, 'A4', 'C7', 'Sipinsur', '3', '2023-07-26 15:02:55.706'),
(83, 'A4', 'C8', 'Sipinsur', '4', '2023-07-26 15:02:55.706'),
(84, 'A5', 'C4', 'Rumah Tjong A Fie', '2', '2023-07-26 15:03:38.495'),
(85, 'A5', 'C2', 'Rumah Tjong A Fie', '3', '2023-07-26 15:03:38.495'),
(86, 'A5', 'C3', 'Rumah Tjong A Fie', '4', '2023-07-26 15:03:38.496'),
(87, 'A5', 'C6', 'Rumah Tjong A Fie', '4', '2023-07-26 15:03:38.497'),
(88, 'A5', 'C5', 'Rumah Tjong A Fie', '2', '2023-07-26 15:03:38.496'),
(89, 'A5', 'C1', 'Rumah Tjong A Fie', '4', '2023-07-26 15:03:38.755'),
(90, 'A5', 'C7', 'Rumah Tjong A Fie', '4', '2023-07-26 15:03:38.888'),
(91, 'A5', 'C8', 'Rumah Tjong A Fie', '2', '2023-07-26 15:03:38.894'),
(92, 'A6', 'C1', 'Danau Lau Kawar', '2', '2023-07-26 15:04:19.650'),
(93, 'A6', 'C2', 'Danau Lau Kawar', '3', '2023-07-26 15:04:19.651'),
(94, 'A6', 'C3', 'Danau Lau Kawar', '2', '2023-07-26 15:04:19.653'),
(95, 'A6', 'C6', 'Danau Lau Kawar', '3', '2023-07-26 15:04:19.654'),
(96, 'A6', 'C4', 'Danau Lau Kawar', '3', '2023-07-26 15:04:19.654'),
(97, 'A6', 'C5', 'Danau Lau Kawar', '2', '2023-07-26 15:04:19.654'),
(98, 'A6', 'C7', 'Danau Lau Kawar', '2', '2023-07-26 15:04:20.048'),
(99, 'A6', 'C8', 'Danau Lau Kawar', '4', '2023-07-26 15:04:20.158'),
(100, 'A7', 'C1', 'Bukit Gundaling', '4', '2023-07-26 15:04:51.013'),
(101, 'A7', 'C2', 'Bukit Gundaling', '4', '2023-07-26 15:04:51.013'),
(102, 'A7', 'C3', 'Bukit Gundaling', '4', '2023-07-26 15:04:51.014'),
(103, 'A7', 'C4', 'Bukit Gundaling', '4', '2023-07-26 15:04:51.014'),
(104, 'A7', 'C6', 'Bukit Gundaling', '2', '2023-07-26 15:04:51.016'),
(105, 'A7', 'C5', 'Bukit Gundaling', '3', '2023-07-26 15:04:51.016'),
(106, 'A7', 'C7', 'Bukit Gundaling', '2', '2023-07-26 15:04:51.127'),
(107, 'A7', 'C8', 'Bukit Gundaling', '4', '2023-07-26 15:04:51.127'),
(108, 'A8', 'C1', 'Funland Mikie Holiday Resort & Hotel', '3', '2023-07-26 15:06:09.219'),
(109, 'A8', 'C2', 'Funland Mikie Holiday Resort & Hotel', '3', '2023-07-26 15:06:09.219'),
(110, 'A8', 'C3', 'Funland Mikie Holiday Resort & Hotel', '4', '2023-07-26 15:06:09.219'),
(111, 'A8', 'C4', 'Funland Mikie Holiday Resort & Hotel', '4', '2023-07-26 15:06:09.219'),
(112, 'A8', 'C6', 'Funland Mikie Holiday Resort & Hotel', '4', '2023-07-26 15:06:09.221'),
(113, 'A8', 'C5', 'Funland Mikie Holiday Resort & Hotel', '2', '2023-07-26 15:06:09.221'),
(114, 'A8', 'C7', 'Funland Mikie Holiday Resort & Hotel', '4', '2023-07-26 15:06:09.623'),
(115, 'A8', 'C8', 'Funland Mikie Holiday Resort & Hotel', '4', '2023-07-26 15:06:09.624'),
(116, 'A9', 'C1', 'green HILL city', '4', '2023-07-26 15:06:38.146'),
(117, 'A9', 'C2', 'green HILL city', '2', '2023-07-26 15:06:38.146'),
(118, 'A9', 'C3', 'green HILL city', '4', '2023-07-26 15:06:38.146'),
(119, 'A9', 'C4', 'green HILL city', '2', '2023-07-26 15:06:38.146'),
(120, 'A9', 'C6', 'green HILL city', '2', '2023-07-26 15:06:38.148'),
(121, 'A9', 'C5', 'green HILL city', '4', '2023-07-26 15:06:38.148'),
(122, 'A9', 'C7', 'green HILL city', '4', '2023-07-26 15:06:38.266'),
(123, 'A9', 'C8', 'green HILL city', '3', '2023-07-26 15:06:38.313'),
(124, 'A10', 'C1', 'Pantai Sorake dan Pantai Lagundri', '3', '2023-07-26 15:07:11.952'),
(125, 'A10', 'C3', 'Pantai Sorake dan Pantai Lagundri', '3', '2023-07-26 15:07:11.952'),
(126, 'A10', 'C4', 'Pantai Sorake dan Pantai Lagundri', '4', '2023-07-26 15:07:11.952'),
(127, 'A10', 'C6', 'Pantai Sorake dan Pantai Lagundri', '2', '2023-07-26 15:07:11.953'),
(128, 'A10', 'C7', 'Pantai Sorake dan Pantai Lagundri', '4', '2023-07-26 15:07:11.953'),
(129, 'A10', 'C5', 'Pantai Sorake dan Pantai Lagundri', '2', '2023-07-26 15:07:11.953'),
(130, 'A10', 'C8', 'Pantai Sorake dan Pantai Lagundri', '3', '2023-07-26 15:07:12.112'),
(131, 'A10', 'C2', 'Pantai Sorake dan Pantai Lagundri', '2', '2023-07-26 15:07:12.156'),
(132, 'A11', 'C1', 'Rahmat International Wildlife Museum and Gallery ', '3', '2023-07-26 15:08:08.286'),
(133, 'A11', 'C3', 'Rahmat International Wildlife Museum and Gallery ', '2', '2023-07-26 15:08:08.287'),
(134, 'A11', 'C4', 'Rahmat International Wildlife Museum and Gallery ', '3', '2023-07-26 15:08:08.288'),
(135, 'A11', 'C5', 'Rahmat International Wildlife Museum and Gallery ', '4', '2023-07-26 15:08:08.288'),
(136, 'A11', 'C6', 'Rahmat International Wildlife Museum and Gallery ', '4', '2023-07-26 15:08:08.289'),
(137, 'A11', 'C7', 'Rahmat International Wildlife Museum and Gallery ', '3', '2023-07-26 15:08:08.289'),
(138, 'A11', 'C2', 'Rahmat International Wildlife Museum and Gallery ', '2', '2023-07-26 15:08:08.672'),
(139, 'A11', 'C8', 'Rahmat International Wildlife Museum and Gallery ', '4', '2023-07-26 15:08:08.698'),
(140, 'A12', 'C1', 'Danau Toba', '4', '2023-07-26 15:08:42.963'),
(141, 'A12', 'C2', 'Danau Toba', '3', '2023-07-26 15:08:42.963'),
(142, 'A12', 'C4', 'Danau Toba', '4', '2023-07-26 15:08:42.964'),
(143, 'A12', 'C3', 'Danau Toba', '3', '2023-07-26 15:08:42.964'),
(144, 'A12', 'C5', 'Danau Toba', '3', '2023-07-26 15:08:42.966'),
(145, 'A12', 'C6', 'Danau Toba', '3', '2023-07-26 15:08:42.966'),
(146, 'A12', 'C7', 'Danau Toba', '3', '2023-07-26 15:08:43.023'),
(147, 'A12', 'C8', 'Danau Toba', '2', '2023-07-26 15:08:43.025'),
(164, 'A13', 'C1', 'Pulau Samosir', '4', '2023-07-26 15:17:35.221'),
(165, 'A13', 'C2', 'Pulau Samosir', '3', '2023-07-26 15:17:35.221'),
(166, 'A13', 'C4', 'Pulau Samosir', '2', '2023-07-26 15:17:35.222'),
(167, 'A13', 'C3', 'Pulau Samosir', '3', '2023-07-26 15:17:35.222'),
(168, 'A13', 'C5', 'Pulau Samosir', '4', '2023-07-26 15:17:35.224'),
(169, 'A13', 'C6', 'Pulau Samosir', '3', '2023-07-26 15:17:35.224'),
(170, 'A13', 'C7', 'Pulau Samosir', '4', '2023-07-26 15:17:35.613'),
(171, 'A13', 'C8', 'Pulau Samosir', '3', '2023-07-26 15:17:35.646'),
(212, 'A14', 'C8', 'Taman Nasional Gunung Leuser', '2', '2023-07-27 21:14:41.409'),
(213, 'A14', 'C1', 'Taman Nasional Gunung Leuser', '3', '2023-07-27 21:14:41.409'),
(214, 'A14', 'C5', 'Taman Nasional Gunung Leuser', '2', '2023-07-27 21:14:41.409'),
(215, 'A14', 'C7', 'Taman Nasional Gunung Leuser', '2', '2023-07-27 21:14:41.409'),
(216, 'A14', 'C3', 'Taman Nasional Gunung Leuser', '4', '2023-07-27 21:14:41.409'),
(217, 'A14', 'C6', 'Taman Nasional Gunung Leuser', '4', '2023-07-27 21:14:41.409'),
(218, 'A14', 'C2', 'Taman Nasional Gunung Leuser', '4', '2023-07-27 21:14:41.836'),
(219, 'A14', 'C4', 'Taman Nasional Gunung Leuser', '4', '2023-07-27 21:14:41.867'),
(220, 'A15', 'C1', 'Air Terjun Sipiso-piso', '3', '2023-07-27 21:15:10.804'),
(221, 'A15', 'C3', 'Air Terjun Sipiso-piso', '2', '2023-07-27 21:15:10.805'),
(222, 'A15', 'C4', 'Air Terjun Sipiso-piso', '4', '2023-07-27 21:15:10.806'),
(223, 'A15', 'C6', 'Air Terjun Sipiso-piso', '2', '2023-07-27 21:15:10.807'),
(224, 'A15', 'C5', 'Air Terjun Sipiso-piso', '4', '2023-07-27 21:15:10.807'),
(225, 'A15', 'C2', 'Air Terjun Sipiso-piso', '2', '2023-07-27 21:15:10.807'),
(226, 'A15', 'C7', 'Air Terjun Sipiso-piso', '2', '2023-07-27 21:15:11.197'),
(227, 'A15', 'C8', 'Air Terjun Sipiso-piso', '3', '2023-07-27 21:15:11.231'),
(228, 'A16', 'C6', 'Salju Panas Tinggi Raja', '3', '2023-07-27 21:15:50.936'),
(229, 'A16', 'C1', 'Salju Panas Tinggi Raja', '3', '2023-07-27 21:15:50.936'),
(230, 'A16', 'C7', 'Salju Panas Tinggi Raja', '2', '2023-07-27 21:15:50.936'),
(231, 'A16', 'C8', 'Salju Panas Tinggi Raja', '3', '2023-07-27 21:15:50.939'),
(232, 'A16', 'C3', 'Salju Panas Tinggi Raja', '4', '2023-07-27 21:15:50.939'),
(233, 'A16', 'C2', 'Salju Panas Tinggi Raja', '3', '2023-07-27 21:15:50.939'),
(234, 'A16', 'C4', 'Salju Panas Tinggi Raja', '2', '2023-07-27 21:15:51.320'),
(235, 'A16', 'C5', 'Salju Panas Tinggi Raja', '3', '2023-07-27 21:15:51.361'),
(236, 'A17', 'C4', 'Mesjid Raya Azizi Tanjung Pura', '3', '2023-07-27 21:16:18.044'),
(237, 'A17', 'C1', 'Mesjid Raya Azizi Tanjung Pura', '3', '2023-07-27 21:16:18.044'),
(238, 'A17', 'C2', 'Mesjid Raya Azizi Tanjung Pura', '4', '2023-07-27 21:16:18.044'),
(239, 'A17', 'C5', 'Mesjid Raya Azizi Tanjung Pura', '3', '2023-07-27 21:16:18.044'),
(240, 'A17', 'C3', 'Mesjid Raya Azizi Tanjung Pura', '3', '2023-07-27 21:16:18.045'),
(241, 'A17', 'C6', 'Mesjid Raya Azizi Tanjung Pura', '4', '2023-07-27 21:16:18.045'),
(242, 'A17', 'C7', 'Mesjid Raya Azizi Tanjung Pura', '3', '2023-07-27 21:16:18.158'),
(243, 'A17', 'C8', 'Mesjid Raya Azizi Tanjung Pura', '3', '2023-07-27 21:16:18.193'),
(244, 'A18', 'C2', 'Kampung Suluk Babussalam Tanjung Pura', '2', '2023-07-27 21:16:42.475'),
(245, 'A18', 'C3', 'Kampung Suluk Babussalam Tanjung Pura', '3', '2023-07-27 21:16:42.475'),
(246, 'A18', 'C1', 'Kampung Suluk Babussalam Tanjung Pura', '2', '2023-07-27 21:16:42.475'),
(247, 'A18', 'C4', 'Kampung Suluk Babussalam Tanjung Pura', '3', '2023-07-27 21:16:42.475'),
(248, 'A18', 'C5', 'Kampung Suluk Babussalam Tanjung Pura', '4', '2023-07-27 21:16:42.478'),
(249, 'A18', 'C6', 'Kampung Suluk Babussalam Tanjung Pura', '4', '2023-07-27 21:16:42.479'),
(250, 'A18', 'C7', 'Kampung Suluk Babussalam Tanjung Pura', '3', '2023-07-27 21:16:42.591'),
(251, 'A18', 'C8', 'Kampung Suluk Babussalam Tanjung Pura', '2', '2023-07-27 21:16:42.626');

-- --------------------------------------------------------

--
-- Table structure for table `kriteria`
--

CREATE TABLE `kriteria` (
  `kode` varchar(255) NOT NULL,
  `nama_kriteria` varchar(255) DEFAULT NULL,
  `bobot` double DEFAULT NULL,
  `kategory` varchar(255) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `kriteria`
--

INSERT INTO `kriteria` (`kode`, `nama_kriteria`, `bobot`, `kategory`, `created_at`) VALUES
('C1', 'Lokasi', 0.05, 'Cost', '2023-07-25 14:42:59.000'),
('C2', 'Kualitas Lingkungan', 0.05, 'Benefit', '2023-07-25 21:18:10.743'),
('C3', 'Infrastruktur', 0.1, 'Benefit', '2023-07-25 21:40:04.521'),
('C4', 'Fasilitas TIK', 0.1, 'Benefit', '2023-07-25 21:40:26.320'),
('C5', 'Aksesbilitas', 0.15, 'Benefit', '2023-07-25 21:40:50.206'),
('C6', 'Kuliner', 0.15, 'Benefit', '2023-07-25 21:41:05.137'),
('C7', 'Lingkungan Sosial', 0.2, 'Benefit', '2023-07-25 21:41:21.682'),
('C8', 'Kebijakan Daerah', 0.2, 'Benefit', '2023-07-25 21:49:40.889');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` bigint(20) NOT NULL,
  `username` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `role` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `username`, `password`, `role`) VALUES
(1, 'admin', '$2a$10$s2EqSIaDG3OBQnbUKIkfOentCIPdbiNxScTxU310Jz8lDcgfIpKq2', 'admin'),
(2, 'admin1', '$2a$10$CfS1JnGhh9Vsrk4FDhemXO1Pq3AKBbizGXSbyq8n3fwlzOokpRJ3K', 'admin'),
(3, 'tamu', '$2a$10$rRGv0LlyoSVoC49AOBHCzeQ9xfrSlfyiCwDQ9U2UfgGZjQN8D7Jgu', 'tamu');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `alternatifs`
--
ALTER TABLE `alternatifs`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_alternatifs_kriteria` (`kode_kriteria`);

--
-- Indexes for table `kriteria`
--
ALTER TABLE `kriteria`
  ADD PRIMARY KEY (`kode`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `alternatifs`
--
ALTER TABLE `alternatifs`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=252;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `alternatifs`
--
ALTER TABLE `alternatifs`
  ADD CONSTRAINT `fk_alternatifs_kriteria` FOREIGN KEY (`kode_kriteria`) REFERENCES `kriteria` (`kode`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
