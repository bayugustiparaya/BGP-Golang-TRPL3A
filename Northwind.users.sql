
--
-- Database: `northwind`
--

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `nama_depan` varchar(50) NULL,
  `nama_belakang` varchar(50) NULL,
  `email` varchar(50) NULL,
  `username` varchar(100) NULL,
  `password` varchar(100) NULL
) ;


