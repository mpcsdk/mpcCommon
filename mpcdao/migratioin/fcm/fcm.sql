--
-- PostgreSQL database dump
--

-- Dumped from database version 12.19 (Ubuntu 12.19-1.pgdg22.04+1)
-- Dumped by pg_dump version 16.3 (Ubuntu 16.3-1.pgdg22.04+1)

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

--
-- Name: public; Type: SCHEMA; Schema: -; Owner: postgres
--

-- *not* creating schema, since initdb creates it


ALTER SCHEMA public OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: fcm_offline_msg; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.fcm_offline_msg (
    fmc_token character varying(255) NOT NULL,
    title character varying(255) NOT NULL,
    body character varying(255) NOT NULL,
    data character varying NOT NULL,
    address character varying(255) NOT NULL,
    user_id character varying(255) NOT NULL,
    created_time timestamp(6) without time zone NOT NULL,
    id character varying NOT NULL
);


ALTER TABLE public.fcm_offline_msg OWNER TO postgres;

--
-- Name: fcm_token; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.fcm_token (
    user_id character varying(255) NOT NULL,
    fcm_token character varying NOT NULL,
    token character varying NOT NULL,
    address character varying(255) NOT NULL,
    created_time timestamp without time zone,
    updated_time timestamp without time zone
);


ALTER TABLE public.fcm_token OWNER TO postgres;

--
-- Name: push_err; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.push_err (
    fmc_token character varying(255),
    title character varying(255),
    body character varying(255),
    data character varying,
    err character varying,
    address character varying(255),
    user_id character varying(255),
    created_time timestamp without time zone
);


ALTER TABLE public.push_err OWNER TO postgres;

--
-- Name: fcm_offline_msg fcm_offline_msg_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.fcm_offline_msg
    ADD CONSTRAINT fcm_offline_msg_pkey PRIMARY KEY (id);


--
-- Name: fcm_token fcm_token_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.fcm_token
    ADD CONSTRAINT fcm_token_pkey PRIMARY KEY (address);


--
-- Name: addrfcm; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX addrfcm ON public.fcm_token USING btree (address, fcm_token, user_id);


--
-- Name: history_msg_address_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX history_msg_address_idx ON public.fcm_offline_msg USING btree (address);


--
-- Name: SCHEMA public; Type: ACL; Schema: -; Owner: postgres
--

REVOKE USAGE ON SCHEMA public FROM PUBLIC;
GRANT ALL ON SCHEMA public TO PUBLIC;


--
-- PostgreSQL database dump complete
--

