PGDMP                         {            gorm_challenge    15.2    15.2     �           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            �           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            �           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            �           1262    16443    gorm_challenge    DATABASE     �   CREATE DATABASE gorm_challenge WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'English_Indonesia.1252';
    DROP DATABASE gorm_challenge;
                postgres    false            �            1259    16445    books    TABLE     �   CREATE TABLE public.books (
    id bigint NOT NULL,
    name_book character varying(100) NOT NULL,
    author character varying(100) NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);
    DROP TABLE public.books;
       public         heap    postgres    false            �            1259    16444    books_id_seq    SEQUENCE     u   CREATE SEQUENCE public.books_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 #   DROP SEQUENCE public.books_id_seq;
       public          postgres    false    215            �           0    0    books_id_seq    SEQUENCE OWNED BY     =   ALTER SEQUENCE public.books_id_seq OWNED BY public.books.id;
          public          postgres    false    214            e           2604    16448    books id    DEFAULT     d   ALTER TABLE ONLY public.books ALTER COLUMN id SET DEFAULT nextval('public.books_id_seq'::regclass);
 7   ALTER TABLE public.books ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    214    215    215            �          0    16445    books 
   TABLE DATA           N   COPY public.books (id, name_book, author, created_at, updated_at) FROM stdin;
    public          postgres    false    215   �
       �           0    0    books_id_seq    SEQUENCE SET     :   SELECT pg_catalog.setval('public.books_id_seq', 3, true);
          public          postgres    false    214            g           2606    16450    books books_pkey 
   CONSTRAINT     N   ALTER TABLE ONLY public.books
    ADD CONSTRAINT books_pkey PRIMARY KEY (id);
 :   ALTER TABLE ONLY public.books DROP CONSTRAINT books_pkey;
       public            postgres    false    215            �   �   x�m��
�0@��+n���]R۬J���kh�B��~�����x������i�x^�:���K��<�E�# z˾dS6�F����:�hU�#�V�w�-1tA^�U�9�?E�M���0=��MUY�������Z ��0X     