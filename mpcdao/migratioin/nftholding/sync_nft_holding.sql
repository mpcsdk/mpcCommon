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
-- Name: nft_holding; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.nft_holding (
    chain_id bigint NOT NULL,
    address character varying(255) NOT NULL,
    contract character varying(255) NOT NULL,
    token_id character varying NOT NULL,
    value bigint NOT NULL,
    block_number bigint NOT NULL,
    updated_at timestamp(6) with time zone NOT NULL,
    kind character varying(255) NOT NULL
);


ALTER TABLE public.nft_holding OWNER TO postgres;

--
-- Name: nft_holding_stat; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.nft_holding_stat (
    chain_id bigint NOT NULL,
    block_number bigint NOT NULL,
    updated_at timestamp(6) with time zone NOT NULL
);


ALTER TABLE public.nft_holding_stat OWNER TO postgres;

--
-- Name: nft_holding_stat stat_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.nft_holding_stat
    ADD CONSTRAINT stat_pkey PRIMARY KEY (chain_id);


--
-- Name: nft_holding_chain_id_address_contract_token_id_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX nft_holding_chain_id_address_contract_token_id_idx ON public.nft_holding USING btree (address, contract, token_id, chain_id);


--
-- Name: SCHEMA public; Type: ACL; Schema: -; Owner: postgres
--

REVOKE USAGE ON SCHEMA public FROM PUBLIC;
GRANT ALL ON SCHEMA public TO PUBLIC;


--
-- PostgreSQL database dump complete
--

