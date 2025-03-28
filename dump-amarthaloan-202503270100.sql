PGDMP                          }            amarthaloan %   14.17 (Ubuntu 14.17-0ubuntu0.22.04.1) %   14.17 (Ubuntu 14.17-0ubuntu0.22.04.1) :    Z           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            [           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            \           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            ]           1262    25099    amarthaloan    DATABASE     `   CREATE DATABASE amarthaloan WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'en_US.UTF-8';
    DROP DATABASE amarthaloan;
                postgres    false            ^           0    0    DATABASE amarthaloan    ACL     +   GRANT ALL ON DATABASE amarthaloan TO refs;
                   postgres    false    3421                        2615    2200    public    SCHEMA        CREATE SCHEMA public;
    DROP SCHEMA public;
                postgres    false            _           0    0    SCHEMA public    COMMENT     6   COMMENT ON SCHEMA public IS 'standard public schema';
                   postgres    false    3            C           1247    25101 
   loan_state    TYPE     k   CREATE TYPE public.loan_state AS ENUM (
    'proposed',
    'approved',
    'invested',
    'disbursed'
);
    DROP TYPE public.loan_state;
       public          refs    false    3            �            1259    25249    borrower    TABLE     �   CREATE TABLE public.borrower (
    borrower_id integer NOT NULL,
    name character varying NOT NULL,
    created_at timestamp without time zone NOT NULL,
    modified_at timestamp without time zone NOT NULL
);
    DROP TABLE public.borrower;
       public         heap    refs    false    3            �            1259    25248    borrower_borrower_id_seq    SEQUENCE     �   CREATE SEQUENCE public.borrower_borrower_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 /   DROP SEQUENCE public.borrower_borrower_id_seq;
       public          refs    false    3    214            `           0    0    borrower_borrower_id_seq    SEQUENCE OWNED BY     U   ALTER SEQUENCE public.borrower_borrower_id_seq OWNED BY public.borrower.borrower_id;
          public          refs    false    213            �            1259    25121    investor    TABLE     �   CREATE TABLE public.investor (
    investor_id integer NOT NULL,
    name character varying NOT NULL,
    email character varying NOT NULL,
    created_at timestamp without time zone NOT NULL,
    modified_at timestamp without time zone NOT NULL
);
    DROP TABLE public.investor;
       public         heap    refs    false    3            �            1259    25120    investor_investor_id_seq    SEQUENCE     �   CREATE SEQUENCE public.investor_investor_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 /   DROP SEQUENCE public.investor_investor_id_seq;
       public          refs    false    210    3            a           0    0    investor_investor_id_seq    SEQUENCE OWNED BY     U   ALTER SEQUENCE public.investor_investor_id_seq OWNED BY public.investor.investor_id;
          public          refs    false    209            �            1259    25337    loans_investment    TABLE     �   CREATE TABLE public.loans_investment (
    loan_investment_id integer NOT NULL,
    loan_id integer NOT NULL,
    investor_id integer NOT NULL,
    invest_amount integer NOT NULL,
    created_at timestamp without time zone NOT NULL
);
 $   DROP TABLE public.loans_investment;
       public         heap    refs    false    3            �            1259    25336 #   loan_investment_loan_investment_seq    SEQUENCE     �   CREATE SEQUENCE public.loan_investment_loan_investment_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 :   DROP SEQUENCE public.loan_investment_loan_investment_seq;
       public          refs    false    3    220            b           0    0 #   loan_investment_loan_investment_seq    SEQUENCE OWNED BY     o   ALTER SEQUENCE public.loan_investment_loan_investment_seq OWNED BY public.loans_investment.loan_investment_id;
          public          refs    false    219            �            1259    25277    loans    TABLE     !  CREATE TABLE public.loans (
    loan_id integer NOT NULL,
    created_by_id integer NOT NULL,
    borrower_id integer NOT NULL,
    principal_amount integer,
    rate integer NOT NULL,
    roi integer NOT NULL,
    state character varying NOT NULL,
    agreement_letter text,
    approved_by integer,
    approval_date timestamp without time zone,
    invested_amount integer,
    disbursed_by integer,
    disbursement_date timestamp without time zone,
    created_at timestamp without time zone,
    modified_at timestamp without time zone
);
    DROP TABLE public.loans;
       public         heap    refs    false    3            �            1259    25322    loans_proof    TABLE     �   CREATE TABLE public.loans_proof (
    loans_proof_id integer NOT NULL,
    loan_id integer NOT NULL,
    staff_id integer NOT NULL,
    proof_picture character varying NOT NULL,
    created_at timestamp without time zone NOT NULL
);
    DROP TABLE public.loans_proof;
       public         heap    refs    false    3            �            1259    25321 $   loans_approval_loans_approval_id_seq    SEQUENCE     �   CREATE SEQUENCE public.loans_approval_loans_approval_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 ;   DROP SEQUENCE public.loans_approval_loans_approval_id_seq;
       public          refs    false    3    218            c           0    0 $   loans_approval_loans_approval_id_seq    SEQUENCE OWNED BY     g   ALTER SEQUENCE public.loans_approval_loans_approval_id_seq OWNED BY public.loans_proof.loans_proof_id;
          public          refs    false    217            �            1259    25374    loans_disburse    TABLE     �   CREATE TABLE public.loans_disburse (
    loans_disburse_id integer NOT NULL,
    loan_id integer NOT NULL,
    staff_id integer NOT NULL,
    signed_agreement text NOT NULL,
    created_at timestamp without time zone NOT NULL
);
 "   DROP TABLE public.loans_disburse;
       public         heap    refs    false    3            �            1259    25373 $   loans_disburse_loans_disburse_id_seq    SEQUENCE     �   CREATE SEQUENCE public.loans_disburse_loans_disburse_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 ;   DROP SEQUENCE public.loans_disburse_loans_disburse_id_seq;
       public          refs    false    222    3            d           0    0 $   loans_disburse_loans_disburse_id_seq    SEQUENCE OWNED BY     m   ALTER SEQUENCE public.loans_disburse_loans_disburse_id_seq OWNED BY public.loans_disburse.loans_disburse_id;
          public          refs    false    221            �            1259    25276    loans_loan_id_seq    SEQUENCE     �   CREATE SEQUENCE public.loans_loan_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 (   DROP SEQUENCE public.loans_loan_id_seq;
       public          refs    false    216    3            e           0    0    loans_loan_id_seq    SEQUENCE OWNED BY     G   ALTER SEQUENCE public.loans_loan_id_seq OWNED BY public.loans.loan_id;
          public          refs    false    215            �            1259    25130    staff    TABLE       CREATE TABLE public.staff (
    staff_id integer NOT NULL,
    name character varying NOT NULL,
    role character varying NOT NULL,
    created_at timestamp without time zone NOT NULL,
    modified_at timestamp without time zone NOT NULL,
    email character varying NOT NULL
);
    DROP TABLE public.staff;
       public         heap    refs    false    3            �            1259    25129    staff_id_staff_id_seq    SEQUENCE     �   CREATE SEQUENCE public.staff_id_staff_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 ,   DROP SEQUENCE public.staff_id_staff_id_seq;
       public          refs    false    3    212            f           0    0    staff_id_staff_id_seq    SEQUENCE OWNED BY     L   ALTER SEQUENCE public.staff_id_staff_id_seq OWNED BY public.staff.staff_id;
          public          refs    false    211            �           2604    25252    borrower borrower_id    DEFAULT     |   ALTER TABLE ONLY public.borrower ALTER COLUMN borrower_id SET DEFAULT nextval('public.borrower_borrower_id_seq'::regclass);
 C   ALTER TABLE public.borrower ALTER COLUMN borrower_id DROP DEFAULT;
       public          refs    false    214    213    214            �           2604    25124    investor investor_id    DEFAULT     |   ALTER TABLE ONLY public.investor ALTER COLUMN investor_id SET DEFAULT nextval('public.investor_investor_id_seq'::regclass);
 C   ALTER TABLE public.investor ALTER COLUMN investor_id DROP DEFAULT;
       public          refs    false    209    210    210            �           2604    25280    loans loan_id    DEFAULT     n   ALTER TABLE ONLY public.loans ALTER COLUMN loan_id SET DEFAULT nextval('public.loans_loan_id_seq'::regclass);
 <   ALTER TABLE public.loans ALTER COLUMN loan_id DROP DEFAULT;
       public          refs    false    215    216    216            �           2604    25377     loans_disburse loans_disburse_id    DEFAULT     �   ALTER TABLE ONLY public.loans_disburse ALTER COLUMN loans_disburse_id SET DEFAULT nextval('public.loans_disburse_loans_disburse_id_seq'::regclass);
 O   ALTER TABLE public.loans_disburse ALTER COLUMN loans_disburse_id DROP DEFAULT;
       public          refs    false    221    222    222            �           2604    25340 #   loans_investment loan_investment_id    DEFAULT     �   ALTER TABLE ONLY public.loans_investment ALTER COLUMN loan_investment_id SET DEFAULT nextval('public.loan_investment_loan_investment_seq'::regclass);
 R   ALTER TABLE public.loans_investment ALTER COLUMN loan_investment_id DROP DEFAULT;
       public          refs    false    219    220    220            �           2604    25325    loans_proof loans_proof_id    DEFAULT     �   ALTER TABLE ONLY public.loans_proof ALTER COLUMN loans_proof_id SET DEFAULT nextval('public.loans_approval_loans_approval_id_seq'::regclass);
 I   ALTER TABLE public.loans_proof ALTER COLUMN loans_proof_id DROP DEFAULT;
       public          refs    false    217    218    218            �           2604    25133    staff staff_id    DEFAULT     s   ALTER TABLE ONLY public.staff ALTER COLUMN staff_id SET DEFAULT nextval('public.staff_id_staff_id_seq'::regclass);
 =   ALTER TABLE public.staff ALTER COLUMN staff_id DROP DEFAULT;
       public          refs    false    212    211    212            O          0    25249    borrower 
   TABLE DATA           N   COPY public.borrower (borrower_id, name, created_at, modified_at) FROM stdin;
    public          refs    false    214   {D       K          0    25121    investor 
   TABLE DATA           U   COPY public.investor (investor_id, name, email, created_at, modified_at) FROM stdin;
    public          refs    false    210   �D       Q          0    25277    loans 
   TABLE DATA           �   COPY public.loans (loan_id, created_by_id, borrower_id, principal_amount, rate, roi, state, agreement_letter, approved_by, approval_date, invested_amount, disbursed_by, disbursement_date, created_at, modified_at) FROM stdin;
    public          refs    false    216   �D       W          0    25374    loans_disburse 
   TABLE DATA           l   COPY public.loans_disburse (loans_disburse_id, loan_id, staff_id, signed_agreement, created_at) FROM stdin;
    public          refs    false    222   �D       U          0    25337    loans_investment 
   TABLE DATA           o   COPY public.loans_investment (loan_investment_id, loan_id, investor_id, invest_amount, created_at) FROM stdin;
    public          refs    false    220   �D       S          0    25322    loans_proof 
   TABLE DATA           c   COPY public.loans_proof (loans_proof_id, loan_id, staff_id, proof_picture, created_at) FROM stdin;
    public          refs    false    218   E       M          0    25130    staff 
   TABLE DATA           U   COPY public.staff (staff_id, name, role, created_at, modified_at, email) FROM stdin;
    public          refs    false    212   )E       g           0    0    borrower_borrower_id_seq    SEQUENCE SET     G   SELECT pg_catalog.setval('public.borrower_borrower_id_seq', 1, false);
          public          refs    false    213            h           0    0    investor_investor_id_seq    SEQUENCE SET     G   SELECT pg_catalog.setval('public.investor_investor_id_seq', 1, false);
          public          refs    false    209            i           0    0 #   loan_investment_loan_investment_seq    SEQUENCE SET     R   SELECT pg_catalog.setval('public.loan_investment_loan_investment_seq', 1, false);
          public          refs    false    219            j           0    0 $   loans_approval_loans_approval_id_seq    SEQUENCE SET     S   SELECT pg_catalog.setval('public.loans_approval_loans_approval_id_seq', 1, false);
          public          refs    false    217            k           0    0 $   loans_disburse_loans_disburse_id_seq    SEQUENCE SET     S   SELECT pg_catalog.setval('public.loans_disburse_loans_disburse_id_seq', 1, false);
          public          refs    false    221            l           0    0    loans_loan_id_seq    SEQUENCE SET     @   SELECT pg_catalog.setval('public.loans_loan_id_seq', 1, false);
          public          refs    false    215            m           0    0    staff_id_staff_id_seq    SEQUENCE SET     D   SELECT pg_catalog.setval('public.staff_id_staff_id_seq', 1, false);
          public          refs    false    211            �           2606    25256    borrower borrower_pk 
   CONSTRAINT     [   ALTER TABLE ONLY public.borrower
    ADD CONSTRAINT borrower_pk PRIMARY KEY (borrower_id);
 >   ALTER TABLE ONLY public.borrower DROP CONSTRAINT borrower_pk;
       public            refs    false    214            �           2606    25128    investor investor_pk 
   CONSTRAINT     [   ALTER TABLE ONLY public.investor
    ADD CONSTRAINT investor_pk PRIMARY KEY (investor_id);
 >   ALTER TABLE ONLY public.investor DROP CONSTRAINT investor_pk;
       public            refs    false    210            �           2606    25342 #   loans_investment loan_investment_pk 
   CONSTRAINT     q   ALTER TABLE ONLY public.loans_investment
    ADD CONSTRAINT loan_investment_pk PRIMARY KEY (loan_investment_id);
 M   ALTER TABLE ONLY public.loans_investment DROP CONSTRAINT loan_investment_pk;
       public            refs    false    220            �           2606    25329    loans_proof loans_approval_pkey 
   CONSTRAINT     i   ALTER TABLE ONLY public.loans_proof
    ADD CONSTRAINT loans_approval_pkey PRIMARY KEY (loans_proof_id);
 I   ALTER TABLE ONLY public.loans_proof DROP CONSTRAINT loans_approval_pkey;
       public            refs    false    218            �           2606    25381     loans_disburse loans_disburse_pk 
   CONSTRAINT     m   ALTER TABLE ONLY public.loans_disburse
    ADD CONSTRAINT loans_disburse_pk PRIMARY KEY (loans_disburse_id);
 J   ALTER TABLE ONLY public.loans_disburse DROP CONSTRAINT loans_disburse_pk;
       public            refs    false    222            �           2606    25284    loans loans_pk 
   CONSTRAINT     Q   ALTER TABLE ONLY public.loans
    ADD CONSTRAINT loans_pk PRIMARY KEY (loan_id);
 8   ALTER TABLE ONLY public.loans DROP CONSTRAINT loans_pk;
       public            refs    false    216            �           2606    25139    staff staff_email 
   CONSTRAINT     M   ALTER TABLE ONLY public.staff
    ADD CONSTRAINT staff_email UNIQUE (email);
 ;   ALTER TABLE ONLY public.staff DROP CONSTRAINT staff_email;
       public            refs    false    212            �           2606    25137    staff staff_id_pk 
   CONSTRAINT     U   ALTER TABLE ONLY public.staff
    ADD CONSTRAINT staff_id_pk PRIMARY KEY (staff_id);
 ;   ALTER TABLE ONLY public.staff DROP CONSTRAINT staff_id_pk;
       public            refs    false    212            O      x������ � �      K      x������ � �      Q      x������ � �      W      x������ � �      U      x������ � �      S      x������ � �      M      x������ � �     