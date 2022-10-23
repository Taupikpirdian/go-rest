-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost:3306
-- Generation Time: Oct 23, 2022 at 11:48 AM
-- Server version: 5.7.33
-- PHP Version: 7.4.19

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `go_rest`
--

-- --------------------------------------------------------

--
-- Table structure for table `students`
--

CREATE TABLE `students` (
  `id` int(11) NOT NULL,
  `nim` varchar(25) NOT NULL,
  `name` varchar(50) NOT NULL,
  `gender` enum('male','female') NOT NULL,
  `dob` date NOT NULL,
  `pob` varchar(25) NOT NULL,
  `jenjang` varchar(5) NOT NULL,
  `study_program` varchar(50) NOT NULL,
  `faculty` varchar(50) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `students`
--

INSERT INTO `students` (`id`, `nim`, `name`, `gender`, `dob`, `pob`, `jenjang`, `study_program`, `faculty`, `created_at`, `updated_at`) VALUES
(5, '41037006151017', 'Taupik Pirdian', 'male', '1996-03-19', 'Bandung', 'S1', 'Teknik Informatika', 'Teknik', '2022-10-22 17:00:00', '2022-10-22 17:00:00');

-- --------------------------------------------------------

--
-- Table structure for table `student_achievements`
--

CREATE TABLE `student_achievements` (
  `id` int(11) NOT NULL,
  `date` date NOT NULL,
  `name` varchar(50) NOT NULL,
  `ranking` varchar(10) NOT NULL,
  `category` varchar(25) NOT NULL,
  `level` varchar(25) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `student_achievement_members`
--

CREATE TABLE `student_achievement_members` (
  `id` int(11) NOT NULL,
  `student_achievement_id` int(11) NOT NULL,
  `student_id` int(11) NOT NULL,
  `is_captain` tinyint(1) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `students`
--
ALTER TABLE `students`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `student_achievements`
--
ALTER TABLE `student_achievements`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `student_achievement_members`
--
ALTER TABLE `student_achievement_members`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `students`
--
ALTER TABLE `students`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT for table `student_achievements`
--
ALTER TABLE `student_achievements`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `student_achievement_members`
--
ALTER TABLE `student_achievement_members`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
