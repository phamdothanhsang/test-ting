CREATE SCHEMA IF NOT EXISTS `kol_test`;

-- Dummies data generated from https://filldb.info/

DROP TABLE IF EXISTS `KOL`;
/*!40101 SET @saved_cs_client= @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `KOL` (
  `KolID` int(11) NOT NULL,
  `UserProfileID` int(11) DEFAULT NULL,
  `Language` text DEFAULT NULL,
  `Education` text DEFAULT NULL,
  `ExpectedSalary` int(11) DEFAULT NULL,
  `ExpectedSalaryEnable` tinyint(1) DEFAULT NULL,
  `ChannelSettingTypeID` int(11) DEFAULT NULL,
  `IDFrontURL` text DEFAULT NULL,
  `IDBackURL` text DEFAULT NULL,
  `PortraitURL` text DEFAULT NULL,
  `RewardID` int(11) DEFAULT NULL,
  `PaymentMethodID` int(11) DEFAULT NULL,
  `TestimonialsID` int(11) DEFAULT NULL,
  `VerificationStatus` tinyint(1) DEFAULT NULL,
  `Enabled` tinyint(1) DEFAULT NULL,
  `ActiveDate` date DEFAULT NULL,
  `Active` tinyint(1) DEFAULT NULL,
  `CreatedBy` text DEFAULT NULL,
  `CreatedDate` date DEFAULT NULL,
  `ModifiedBy` text DEFAULT NULL,
  `ModifiedDate` date DEFAULT NULL,
  `IsRemove` tinyint(1) DEFAULT NULL,
  `IsOnBoarding` tinyint(1) DEFAULT NULL,
  `Code` text DEFAULT NULL,
  `PortraitRightURL` text DEFAULT NULL,
  `PortraitLeftURL` text DEFAULT NULL,
  `LivenessStatus` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`KolID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `KOL`
--

LOCK TABLES `KOL` WRITE;
/*!40000 ALTER TABLE `KOL` DISABLE KEYS */;
INSERT INTO `KOL` VALUES
(2,196989754,'gljw','omnis',372,0,15497423,'http://corwin.net/','http://www.parker.com/','http://www.cremincartwright.com/',8070487,59666948,44783,0,0,'1970-05-17',1,'nchj','1978-11-25','covq','1989-03-14',0,0,'lgul','http://www.hartmann.com/','http://turner.com/',1),
(5,34680,'gemt','veritatis',469260108,1,287,'http://koch.net/','http://www.hodkiewiczyost.com/','http://miller.com/',79,4501197,6,0,0,'2010-02-07',0,'uqst','1996-09-13','bhtg','2011-01-07',0,0,'vylr','http://bashirian.com/','http://www.kunze.com/',0),
(6,3194,'zhax','voluptas',564045,1,131272,'http://www.fritschhowe.com/','http://www.bartellstamm.net/','http://www.armstrong.biz/',0,606194,95396,1,1,'1996-06-21',0,'fbup','1997-05-20','sosm','1990-06-10',1,1,'gato','http://lebsack.com/','http://gusikowskisteuber.com/',0),
(21,60,'dmpf','quia',957530617,0,8280260,'http://hayes.org/','http://weberking.com/','http://www.shields.com/',636,450230,4,0,0,'1973-01-13',1,'bran','2023-02-07','cdhy','1990-02-04',1,0,'qcsn','http://keeling.com/','http://www.nitzscheschiller.com/',1),
(24,71792,'sile','eum',8022,0,8181613,'http://www.bashirian.com/','http://www.rueckerbernier.com/','http://www.roob.com/',9,317731,807288,1,1,'1986-04-14',0,'umvc','2007-09-06','bahm','2023-07-07',0,1,'blsk','http://dickens.com/','http://feil.com/',0),
(82,69,'qsyt','omnis',7918031,1,2,'http://wunsch.com/','http://boyle.org/','http://zulaufbuckridge.com/',247336,5,584452606,1,1,'2021-09-11',1,'zdqu','2005-12-05','ihmw','2004-10-31',1,0,'fcxk','http://www.wolf.com/','http://walker.com/',1),
(3659,6558778,'ctkp','eos',46,1,50687977,'http://becker.net/','http://www.swaniawskiswaniawski.biz/','http://cristthiel.com/',828,95851867,843092367,1,0,'2022-04-28',1,'ojtk','2015-06-21','hiep','1994-01-01',1,1,'gwui','http://braun.com/','http://www.nikolaus.com/',1),
(5391,18611,'erug','iste',7298,1,759,'http://www.nader.com/','http://wolftremblay.org/','http://www.gerhold.com/',8400507,8824,0,0,1,'1995-10-29',0,'hjsj','2003-08-19','taza','2005-03-12',0,1,'vkbw','http://www.eichmannpacocha.com/','http://reilly.com/',0),
(7021,809303,'isxp','nam',0,1,711,'http://www.altenwerth.biz/','http://danielgerhold.info/','http://harveystanton.net/',68,785299780,0,1,0,'2020-01-23',0,'pwvb','1987-04-09','gdcj','1987-12-10',1,0,'myof','http://torp.org/','http://www.bode.biz/',1),
(13045,0,'qywp','cupiditate',6280037,1,0,'http://www.leffler.com/','http://wilkinson.biz/','http://www.murray.com/',195,403836202,4289,1,1,'1988-03-09',1,'uzse','1986-08-27','weci','1996-05-25',0,1,'fmzm','http://davisullrich.com/','http://www.schoen.com/',0),
(26884,13,'qaqu','animi',5649451,0,32,'http://becker.net/','http://www.connellymetz.com/','http://champlin.com/',58946588,0,51459400,1,0,'1992-10-12',0,'besn','1979-02-20','zfle','1982-12-24',0,0,'kutx','http://www.sauer.com/','http://www.feilhane.org/',1),
(48500,83575467,'tdcj','voluptatem',11222,0,956393934,'http://treutel.com/','http://morissettegibson.com/','http://www.kirlin.com/',9711325,485510,0,0,0,'2023-08-31',0,'hxwy','2003-06-12','wofd','1985-09-13',0,0,'szag','http://www.yostoreilly.com/','http://www.gislason.com/',1),
(77461,721917594,'vsyq','doloremque',926428770,0,4,'http://fritsch.com/','http://www.nitzschemayert.com/','http://www.tillman.org/',30595,7533143,2285590,1,0,'2015-06-01',1,'pkeq','1994-01-13','ctet','1974-05-19',1,0,'stgu','http://deckow.info/','http://mccullough.com/',1),
(279091,97122026,'ftge','omnis',486,1,332831143,'http://rauklocko.com/','http://www.schneider.com/','http://www.oberbrunner.com/',4689170,2,188313,0,0,'2003-05-25',1,'kcqe','2022-01-26','wskt','2019-01-22',1,0,'znyh','http://www.wittinglebsack.com/','http://www.mante.net/',0),
(665473,195028099,'gbkm','iure',74172,0,9759,'http://metz.com/','http://willms.biz/','http://www.bartoncronin.com/',56,26736,314885,0,0,'1976-10-13',0,'mjsh','1979-05-30','htkv','1993-01-18',1,1,'hsto','http://www.roob.net/','http://www.cummerata.net/',0),
(986512,457209,'dszx','et',8610,1,4506377,'http://www.lowe.net/','http://beattywintheiser.com/','http://bruen.info/',38727328,3,2958344,1,0,'1982-10-22',0,'rbbo','1988-10-12','hcfp','2022-06-23',0,1,'eajp','http://www.rowe.biz/','http://nienow.com/',1),
(18187187,16217,'zkep','est',4818013,0,53656,'http://www.stokes.net/','http://howellhilpert.biz/','http://www.ratkekuhn.com/',84,98257,3323238,0,1,'1978-02-04',1,'tjxy','2000-06-18','ldwf','1991-02-23',0,0,'lelb','http://www.hudsonmurazik.com/','http://bahringer.info/',0),
(66484519,502,'mqgt','vitae',9891,1,26881,'http://www.bahringerhowe.com/','http://schimmel.com/','http://mraz.org/',4511,92,78919318,1,0,'2000-09-01',0,'gwse','1996-05-09','pbay','2001-05-14',0,0,'llbf','http://lakin.com/','http://www.wintheiser.com/',0),
(111392938,5,'mlmh','neque',21069,0,264036007,'http://www.koepp.com/','http://dubuquejohnson.com/','http://wolf.com/',2671953,6845,800950190,1,0,'2016-12-08',0,'bber','1983-05-22','lkex','2001-04-16',1,1,'qcuf','http://shields.com/','http://medhurststracke.com/',1),
(271652571,7,'mzsz','necessitatibus',65,1,33,'http://www.swaniawski.com/','http://gusikowski.org/','http://predovickautzer.biz/',86627,31,81737358,1,0,'1971-05-15',1,'vmba','1989-07-26','qlrd','1975-01-07',0,1,'huca','http://vonruedenbreitenberg.com/','http://www.ratke.org/',0);
/*!40000 ALTER TABLE `KOL` ENABLE KEYS */;
UNLOCK TABLES;

-- Dump completed on 2024-10-09 13:25:26

