--
-- PostgreSQL database dump
--

-- Dumped from database version 11.5 (Ubuntu 11.5-3.pgdg18.04+1)
-- Dumped by pg_dump version 14.2

-- Started on 2022-04-22 16:09:51 MSK

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

--
-- TOC entry 222 (class 1259 OID 4250527)
-- Name: basketball; Type: TABLE; Schema: public; Owner: mmxxboti
--

CREATE TABLE public.basketball (
    field_0 integer,
    player_name text,
    team_abbreviation text,
    age double precision,
    player_height double precision,
    player_weight double precision,
    college text,
    country text,
    draft_year text,
    draft_round text,
    draft_number text,
    gp integer,
    pts double precision,
    reb double precision,
    ast double precision,
    net_rating double precision,
    oreb_pct double precision,
    dreb_pct double precision,
    usg_pct double precision,
    ts_pct double precision,
    ast_pct double precision,
    season text
);
