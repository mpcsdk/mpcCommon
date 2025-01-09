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
-- Name: chain_tx; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.chain_tx (
    chain_id bigint NOT NULL,
    height bigint NOT NULL,
    block_hash character varying(255) NOT NULL,
    ts bigint NOT NULL,
    tx_hash character varying(255) NOT NULL,
    tx_idx integer NOT NULL,
    log_idx integer NOT NULL,
    "from" character varying(255) NOT NULL,
    "to" character varying(255) NOT NULL,
    contract character varying(255) NOT NULL,
    value character varying(255) NOT NULL,
    gas character varying(255) NOT NULL,
    gas_price character varying(255) NOT NULL,
    nonce bigint NOT NULL,
    kind character varying(255) NOT NULL,
    token_id character varying(255) NOT NULL,
    removed boolean NOT NULL,
    status bigint NOT NULL,
    "traceTag" character varying(255) NOT NULL
);


ALTER TABLE public.chain_tx OWNER TO postgres;

--
-- Name: chain_tx_contract_ts; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX chain_tx_contract_ts ON public.chain_tx USING btree (contract, ts DESC NULLS LAST);


--
-- Name: chain_tx_from_kind_contract_ts_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX chain_tx_from_kind_contract_ts_idx ON public.chain_tx USING btree ("from", kind, contract, ts DESC NULLS LAST);


--
-- Name: chain_tx_height_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX chain_tx_height_idx ON public.chain_tx USING btree (height);


--
-- Name: chain_tx_to_kind_contract_ts_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX chain_tx_to_kind_contract_ts_idx ON public.chain_tx USING btree ("to", kind, contract, ts DESC NULLS LAST);


--
-- Name: chain_tx_ts_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX chain_tx_ts_idx ON public.chain_tx USING btree (ts);


--
-- Name: chain_tx_tx_hash_tx_idx_log_idx_traceTag_token_id_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX "chain_tx_tx_hash_tx_idx_log_idx_traceTag_token_id_idx" ON public.chain_tx USING btree (tx_hash, tx_idx, log_idx, "traceTag");


--
-- Name: SCHEMA public; Type: ACL; Schema: -; Owner: postgres
--

REVOKE USAGE ON SCHEMA public FROM PUBLIC;
GRANT ALL ON SCHEMA public TO PUBLIC;


--
-- PostgreSQL database dump complete
--

