CREATE DATABASE /*!32312 IF NOT EXISTS*/`northwind` /*!40100 DEFAULT CHARACTER SET latin1 */;

USE `northwind`;

--
-- Table structure for table `trans`
--

CREATE TABLE `trans` (
  `id` int(11) AUTO_INCREMENT PRIMARY KEY,
  `merchant_id` varchar(50) NOT NULL,
  `merchant` varchar(50) NOT NULL,
  `code` varchar(50) NOT NULL,
  `name` varchar(50) NOT NULL,
  `signature` varchar(50) NOT NULL
) ENGINE=MyISAM AUTO_INCREMENT=10 DEFAULT CHARSET=utf8;

--
-- Dumping data for table `trans`
--

INSERT INTO `trans` (`merchant_id`, `merchant`, `code`, `name`, `signature`) VALUES 
('98765', 'PARAYA STORE', '707', 'ALFAGROUP', 'bgp0708'),
('98765', 'PARAYA STORE', '409', 'ATMB', 'bgp0708'),
('98765', 'PARAYA STORE', '308', 'BBM Money', 'bgp0708'),
('98765', 'PARAYA STORE', '801', 'BRI Virtual Account', 'bgp0708'),
('98765', 'PARAYA STORE', '500', 'CREDIT CARD', 'bgp0708'),
('98765', 'PARAYA STORE', '701', 'DANAMON ONLINE BANKING', 'bgp0708'),
('98765', 'PARAYA STORE', '708', 'Danamon VA', 'bgp0708'),
('98765', 'PARAYA STORE', '307', 'DOMPETKU', 'bgp0708'),
('98765', 'PARAYA STORE', '802', 'Mandiri Virtual Account', 'bgp0708'),
('98765', 'PARAYA STORE', '402', 'Permata', 'bgp0708'),
('98765', 'PARAYA STORE', '410', 'UNICount', 'bgp0708'),
('98766', 'BGP STORE', '337', 'GOPAY', 'bgp070820'),
('98766', 'BGP STORE', '832', 'OVO', 'bgp070820'),
('98766', 'BGP STORE', '432', 'DANA', 'bgp070820');

-- --------------------------------------------------------
