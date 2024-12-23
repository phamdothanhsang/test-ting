--
-- PostgreSQL database dump
--

-- Dumped from database version 17.2
-- Dumped by pg_dump version 17.2

-- Started on 2024-12-23 16:24:59

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 217 (class 1259 OID 16448)
-- Name: KOL; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."KOL" (
    "KolID" bigint NOT NULL,
    "UserProfileID" bigint,
    "Language" character varying(255),
    "Education" character varying(255),
    "ExpectedSalary" bigint,
    "ExpectedSalaryEnable" boolean,
    "ChannelSettingTypeID" bigint,
    "IDFrontURL" character varying(255),
    "IDBackURL" character varying(255),
    "PortraitURL" character varying(255),
    "RewardID" bigint,
    "PaymentMethodID" bigint,
    "TestimonialsID" bigint,
    "VerificationStatus" boolean,
    "Enabled" boolean,
    "ActiveDate" timestamp without time zone,
    "Active" boolean,
    "CreatedBy" character varying(255),
    "CreatedDate" timestamp without time zone,
    "ModifiedBy" character varying(255),
    "ModifiedDate" timestamp without time zone,
    "IsRemove" boolean,
    "IsOnBoarding" boolean,
    "Code" character varying(255),
    "PortraitRightURL" character varying(255),
    "PortraitLeftURL" character varying(255),
    "LivenessStatus" boolean
);


ALTER TABLE public."KOL" OWNER TO postgres;

--
-- TOC entry 4787 (class 0 OID 16448)
-- Dependencies: 217
-- Data for Name: KOL; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."KOL" ("KolID", "UserProfileID", "Language", "Education", "ExpectedSalary", "ExpectedSalaryEnable", "ChannelSettingTypeID", "IDFrontURL", "IDBackURL", "PortraitURL", "RewardID", "PaymentMethodID", "TestimonialsID", "VerificationStatus", "Enabled", "ActiveDate", "Active", "CreatedBy", "CreatedDate", "ModifiedBy", "ModifiedDate", "IsRemove", "IsOnBoarding", "Code", "PortraitRightURL", "PortraitLeftURL", "LivenessStatus") FROM stdin;
101	2001	en	Bachelor's in Computer Science	70000	t	1	https://example.com/id-front.jpg	https://example.com/id-back.jpg	https://example.com/portrait.jpg	301	401	501	t	t	2024-08-01 09:00:00	t	admin	2024-08-01 08:30:00	admin	2024-08-10 10:15:00	f	t	KOL2024001	https://example.com/portrait-right.jpg	https://example.com/portrait-left.jpg	t
102	2002	vn	Bachelor's in Marketing	50000	f	2	https://example.com/id-front-2.jpg	https://example.com/id-back-2.jpg	https://example.com/portrait-2.jpg	302	402	502	f	t	2024-08-02 09:00:00	t	admin	2024-08-02 08:30:00	admin	2024-08-11 10:15:00	f	f	KOL2024002	https://example.com/portrait-right-2.jpg	https://example.com/portrait-left-2.jpg	f
103	2003	en	Master's in Business Administration	80000	t	1	https://example.com/id-front-3.jpg	https://example.com/id-back-3.jpg	https://example.com/portrait-3.jpg	303	403	503	t	t	2024-08-03 09:00:00	t	admin	2024-08-03 08:30:00	admin	2024-08-12 10:15:00	f	t	KOL2024003	https://example.com/portrait-right-3.jpg	https://example.com/portrait-left-3.jpg	t
104	2004	vn	PhD in Engineering	90000	t	2	https://example.com/id-front-4.jpg	https://example.com/id-back-4.jpg	https://example.com/portrait-4.jpg	304	404	504	t	t	2024-08-04 09:00:00	t	admin	2024-08-04 08:30:00	admin	2024-08-13 10:15:00	f	f	KOL2024004	https://example.com/portrait-right-4.jpg	https://example.com/portrait-left-4.jpg	t
105	2005	en	Bachelor's in Graphic Design	60000	f	1	https://example.com/id-front-5.jpg	https://example.com/id-back-5.jpg	https://example.com/portrait-5.jpg	305	405	505	f	t	2024-08-05 09:00:00	t	admin	2024-08-05 08:30:00	admin	2024-08-14 10:15:00	f	t	KOL2024005	https://example.com/portrait-right-5.jpg	https://example.com/portrait-left-5.jpg	f
106	2006	vn	Diploma in Digital Marketing	40000	t	2	https://example.com/id-front-6.jpg	https://example.com/id-back-6.jpg	https://example.com/portrait-6.jpg	306	406	506	t	t	2024-08-06 09:00:00	t	admin	2024-08-06 08:30:00	admin	2024-08-15 10:15:00	f	t	KOL2024006	https://example.com/portrait-right-6.jpg	https://example.com/portrait-left-6.jpg	t
107	2007	en	Master's in Marketing	75000	f	1	https://example.com/id-front-7.jpg	https://example.com/id-back-7.jpg	https://example.com/portrait-7.jpg	307	407	507	t	t	2024-08-07 09:00:00	t	admin	2024-08-07 08:30:00	admin	2024-08-16 10:15:00	f	t	KOL2024007	https://example.com/portrait-right-7.jpg	https://example.com/portrait-left-7.jpg	t
108	2008	vn	Bachelor's in Communication	65000	t	2	https://example.com/id-front-8.jpg	https://example.com/id-back-8.jpg	https://example.com/portrait-8.jpg	308	408	508	f	t	2024-08-08 09:00:00	t	admin	2024-08-08 08:30:00	admin	2024-08-17 10:15:00	f	f	KOL2024008	https://example.com/portrait-right-8.jpg	https://example.com/portrait-left-8.jpg	f
109	2009	en	Bachelor's in Arts	70000	t	1	https://example.com/id-front-9.jpg	https://example.com/id-back-9.jpg	https://example.com/portrait-9.jpg	309	409	509	t	t	2024-08-09 09:00:00	t	admin	2024-08-09 08:30:00	admin	2024-08-18 10:15:00	f	t	KOL2024009	https://example.com/portrait-right-9.jpg	https://example.com/portrait-left-9.jpg	t
110	2010	vn	Master's in Psychology	85000	f	2	https://example.com/id-front-10.jpg	https://example.com/id-back-10.jpg	https://example.com/portrait-10.jpg	310	410	510	t	t	2024-08-10 09:00:00	t	admin	2024-08-10 08:30:00	admin	2024-08-19 10:15:00	f	f	KOL2024010	https://example.com/portrait-right-10.jpg	https://example.com/portrait-left-10.jpg	t
111	2011	vn	Bachelor's in Information Technology	80000	t	1	https://example.com/id-front-11.jpg	https://example.com/id-back-11.jpg	https://example.com/portrait-11.jpg	311	411	511	t	t	2024-08-11 09:00:00	t	admin	2024-08-11 08:30:00	admin	2024-08-20 10:15:00	f	t	KOL2024011	https://example.com/portrait-right-11.jpg	https://example.com/portrait-left-11.jpg	t
112	2012	en	Master's in Finance	95000	t	2	https://example.com/id-front-12.jpg	https://example.com/id-back-12.jpg	https://example.com/portrait-12.jpg	312	412	512	f	t	2024-08-12 09:00:00	t	admin	2024-08-12 08:30:00	admin	2024-08-21 10:15:00	f	f	KOL2024012	https://example.com/portrait-right-12.jpg	https://example.com/portrait-left-12.jpg	f
113	2013	vn	Bachelor's in Business Administration	65000	f	1	https://example.com/id-front-13.jpg	https://example.com/id-back-13.jpg	https://example.com/portrait-13.jpg	313	413	513	t	t	2024-08-13 09:00:00	t	admin	2024-08-13 08:30:00	admin	2024-08-22 10:15:00	f	t	KOL2024013	https://example.com/portrait-right-13.jpg	https://example.com/portrait-left-13.jpg	t
114	2014	en	Bachelor's in Data Science	85000	t	2	https://example.com/id-front-14.jpg	https://example.com/id-back-14.jpg	https://example.com/portrait-14.jpg	314	414	514	f	t	2024-08-14 09:00:00	t	admin	2024-08-14 08:30:00	admin	2024-08-23 10:15:00	f	f	KOL2024014	https://example.com/portrait-right-14.jpg	https://example.com/portrait-left-14.jpg	f
115	2015	vn	Master's in Software Engineering	90000	t	1	https://example.com/id-front-15.jpg	https://example.com/id-back-15.jpg	https://example.com/portrait-15.jpg	315	415	515	t	t	2024-08-15 09:00:00	t	admin	2024-08-15 08:30:00	admin	2024-08-24 10:15:00	f	t	KOL2024015	https://example.com/portrait-right-15.jpg	https://example.com/portrait-left-15.jpg	t
116	2016	en	PhD in Physics	120000	t	2	https://example.com/id-front-16.jpg	https://example.com/id-back-16.jpg	https://example.com/portrait-16.jpg	316	416	516	t	t	2024-08-16 09:00:00	t	admin	2024-08-16 08:30:00	admin	2024-08-25 10:15:00	f	f	KOL2024016	https://example.com/portrait-right-16.jpg	https://example.com/portrait-left-16.jpg	t
117	2017	vn	Master's in Marketing Management	95000	f	1	https://example.com/id-front-17.jpg	https://example.com/id-back-17.jpg	https://example.com/portrait-17.jpg	317	417	517	t	t	2024-08-17 09:00:00	t	admin	2024-08-17 08:30:00	admin	2024-08-26 10:15:00	f	t	KOL2024017	https://example.com/portrait-right-17.jpg	https://example.com/portrait-left-17.jpg	t
118	2018	en	Bachelor's in Journalism	55000	t	2	https://example.com/id-front-18.jpg	https://example.com/id-back-18.jpg	https://example.com/portrait-18.jpg	318	418	518	t	t	2024-08-18 09:00:00	t	admin	2024-08-18 08:30:00	admin	2024-08-27 10:15:00	f	t	KOL2024018	https://example.com/portrait-right-18.jpg	https://example.com/portrait-left-18.jpg	t
119	2019	vn	Bachelor's in Finance	70000	f	1	https://example.com/id-front-19.jpg	https://example.com/id-back-19.jpg	https://example.com/portrait-19.jpg	319	419	519	f	t	2024-08-19 09:00:00	t	admin	2024-08-19 08:30:00	admin	2024-08-28 10:15:00	f	f	KOL2024019	https://example.com/portrait-right-19.jpg	https://example.com/portrait-left-19.jpg	f
120	2020	en	PhD in Artificial Intelligence	130000	t	2	https://example.com/id-front-20.jpg	https://example.com/id-back-20.jpg	https://example.com/portrait-20.jpg	320	420	520	t	t	2024-08-20 09:00:00	t	admin	2024-08-20 08:30:00	admin	2024-08-29 10:15:00	f	t	KOL2024020	https://example.com/portrait-right-20.jpg	https://example.com/portrait-left-20.jpg	t
121	2021	vn	Master's in Data Science	110000	t	1	https://example.com/id-front-21.jpg	https://example.com/id-back-21.jpg	https://example.com/portrait-21.jpg	321	421	521	f	t	2024-08-21 09:00:00	t	admin	2024-08-21 08:30:00	admin	2024-08-30 10:15:00	f	t	KOL2024021	https://example.com/portrait-right-21.jpg	https://example.com/portrait-left-21.jpg	t
122	2022	en	Bachelor's in Engineering	85000	t	2	https://example.com/id-front-22.jpg	https://example.com/id-back-22.jpg	https://example.com/portrait-22.jpg	322	422	522	t	t	2024-08-22 09:00:00	t	admin	2024-08-22 08:30:00	admin	2024-08-31 10:15:00	f	f	KOL2024022	https://example.com/portrait-right-22.jpg	https://example.com/portrait-left-22.jpg	f
123	2023	vn	Master's in Computer Engineering	95000	t	1	https://example.com/id-front-23.jpg	https://example.com/id-back-23.jpg	https://example.com/portrait-23.jpg	323	423	523	f	t	2024-08-23 09:00:00	t	admin	2024-08-23 08:30:00	admin	2024-09-01 10:15:00	f	t	KOL2024023	https://example.com/portrait-right-23.jpg	https://example.com/portrait-left-23.jpg	t
124	2024	en	Master's in Electronics	105000	t	2	https://example.com/id-front-24.jpg	https://example.com/id-back-24.jpg	https://example.com/portrait-24.jpg	324	424	524	t	t	2024-08-24 09:00:00	t	admin	2024-08-24 08:30:00	admin	2024-09-02 10:15:00	f	f	KOL2024024	https://example.com/portrait-right-24.jpg	https://example.com/portrait-left-24.jpg	f
125	2025	vn	Bachelor's in Electrical Engineering	75000	t	1	https://example.com/id-front-25.jpg	https://example.com/id-back-25.jpg	https://example.com/portrait-25.jpg	325	425	525	t	t	2024-08-25 09:00:00	t	admin	2024-08-25 08:30:00	admin	2024-09-03 10:15:00	f	t	KOL2024025	https://example.com/portrait-right-25.jpg	https://example.com/portrait-left-25.jpg	t
126	2026	en	PhD in Mathematics	115000	f	2	https://example.com/id-front-26.jpg	https://example.com/id-back-26.jpg	https://example.com/portrait-26.jpg	326	426	526	t	t	2024-08-26 09:00:00	t	admin	2024-08-26 08:30:00	admin	2024-09-04 10:15:00	f	f	KOL2024026	https://example.com/portrait-right-26.jpg	https://example.com/portrait-left-26.jpg	t
127	2027	vn	Master's in Business Analytics	100000	t	1	https://example.com/id-front-27.jpg	https://example.com/id-back-27.jpg	https://example.com/portrait-27.jpg	327	427	527	t	t	2024-08-27 09:00:00	t	admin	2024-08-27 08:30:00	admin	2024-09-05 10:15:00	f	t	KOL2024027	https://example.com/portrait-right-27.jpg	https://example.com/portrait-left-27.jpg	f
128	2028	en	Bachelor's in Architecture	70000	f	2	https://example.com/id-front-28.jpg	https://example.com/id-back-28.jpg	https://example.com/portrait-28.jpg	328	428	528	f	t	2024-08-28 09:00:00	t	admin	2024-08-28 08:30:00	admin	2024-09-06 10:15:00	f	t	KOL2024028	https://example.com/portrait-right-28.jpg	https://example.com/portrait-left-28.jpg	t
\.


--
-- TOC entry 4641 (class 2606 OID 16454)
-- Name: KOL KOL_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."KOL"
    ADD CONSTRAINT "KOL_pkey" PRIMARY KEY ("KolID");


-- Completed on 2024-12-23 16:24:59

--
-- PostgreSQL database dump complete
--

