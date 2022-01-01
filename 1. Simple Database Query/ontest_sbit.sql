--
-- Create Database `ontest_sbit` and Use database
--
CREATE DATABASE ontest_sbit;

USE ontest_sbit;

--
-- Create Table `user`
--
CREATE TABLE `user` (
  `ID` int(11) NOT NULL,
  `UserName` varchar(50) NOT NULL,
  `Parent` int(11) NOT NULL
);

--
-- Dumping data for table `user`
--

INSERT INTO `user` (`ID`, `UserName`, `Parent`) VALUES
(1, 'Ali', 2),
(2, 'Budi', 0),
(3, 'Cecep', 1);

--
-- Indexes for table `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`ID`);


SELECT A.ID, A.UserName, B.UserName AS ParentUserName FROM user A LEFT JOIN user B ON A.Parent=B.ID ORDER BY A.ID;