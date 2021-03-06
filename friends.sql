PGDMP     :                	    x            FriendsManagement    13.0    13.0     �           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            �           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            �           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            �           1262    16426    FriendsManagement    DATABASE     w   CREATE DATABASE "FriendsManagement" WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'English_United States.1252';
 #   DROP DATABASE "FriendsManagement";
                postgres    false            �            1259    16443    friends    TABLE     �   CREATE TABLE public.friends (
    id integer NOT NULL,
    user_one_email text NOT NULL,
    user_two_email text NOT NULL,
    update_status boolean NOT NULL
);
    DROP TABLE public.friends;
       public         heap    postgres    false            �            1259    16441    friends_id_seq    SEQUENCE     �   CREATE SEQUENCE public.friends_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 %   DROP SEQUENCE public.friends_id_seq;
       public          postgres    false    201            �           0    0    friends_id_seq    SEQUENCE OWNED BY     A   ALTER SEQUENCE public.friends_id_seq OWNED BY public.friends.id;
          public          postgres    false    200            �            1259    16465    friendships    TABLE       CREATE TABLE public.friendships (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    first_user text NOT NULL,
    second_user text NOT NULL,
    update_status boolean
);
    DROP TABLE public.friendships;
       public         heap    postgres    false            �            1259    16463    friendships_id_seq    SEQUENCE     {   CREATE SEQUENCE public.friendships_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 )   DROP SEQUENCE public.friendships_id_seq;
       public          postgres    false    203            �           0    0    friendships_id_seq    SEQUENCE OWNED BY     I   ALTER SEQUENCE public.friendships_id_seq OWNED BY public.friendships.id;
          public          postgres    false    202            *           2604    16446 
   friends id    DEFAULT     h   ALTER TABLE ONLY public.friends ALTER COLUMN id SET DEFAULT nextval('public.friends_id_seq'::regclass);
 9   ALTER TABLE public.friends ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    200    201    201            +           2604    16468    friendships id    DEFAULT     p   ALTER TABLE ONLY public.friendships ALTER COLUMN id SET DEFAULT nextval('public.friendships_id_seq'::regclass);
 =   ALTER TABLE public.friendships ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    203    202    203            �          0    16443    friends 
   TABLE DATA           T   COPY public.friends (id, user_one_email, user_two_email, update_status) FROM stdin;
    public          postgres    false    201   �       �          0    16465    friendships 
   TABLE DATA           u   COPY public.friendships (id, created_at, updated_at, deleted_at, first_user, second_user, update_status) FROM stdin;
    public          postgres    false    203   �       �           0    0    friends_id_seq    SEQUENCE SET     >   SELECT pg_catalog.setval('public.friends_id_seq', 102, true);
          public          postgres    false    200            �           0    0    friendships_id_seq    SEQUENCE SET     @   SELECT pg_catalog.setval('public.friendships_id_seq', 3, true);
          public          postgres    false    202            -           2606    16451    friends FriendShip_pkey 
   CONSTRAINT     w   ALTER TABLE ONLY public.friends
    ADD CONSTRAINT "FriendShip_pkey" PRIMARY KEY (id, user_one_email, user_two_email);
 C   ALTER TABLE ONLY public.friends DROP CONSTRAINT "FriendShip_pkey";
       public            postgres    false    201    201    201            /           2606    16473    friendships friendships_pkey 
   CONSTRAINT     s   ALTER TABLE ONLY public.friendships
    ADD CONSTRAINT friendships_pkey PRIMARY KEY (id, first_user, second_user);
 F   ALTER TABLE ONLY public.friendships DROP CONSTRAINT friendships_pkey;
       public            postgres    false    203    203    203            0           1259    16474    idx_friendships_deleted_at    INDEX     X   CREATE INDEX idx_friendships_deleted_at ON public.friendships USING btree (deleted_at);
 .   DROP INDEX public.idx_friendships_deleted_at;
       public            postgres    false    203            �   W	  x��Xێܸ}�|1}�~��N ��`������a�K["���ԅ�f�p�$�u9u���9��r��qJ)�q�ag�{��R��{�l�q�F�!w��4�!$������_�1��qrs���>�B�������{�8���4e��m3���r4a�]�À]�q�lo��d[��txߌ�֏��KO'�؅9�`����`��ݛCk��f^�8�"vp/a���Ρ���lusj�ᖃg� ߸OC�;����{��{��'z,��bz��� ۜ��4�	��Gb����A|A@쐇Ʋ.!r�5�n�7����r��[�c�Gm6���4���G�tb�l���p�����9��O)��l͟�k��l���oR�:��:~��"�]�7�Fj�3�kɸ��;�qB���'�8ؗ�oFፃ����T��טOXΨ�	ߧH@����tC���9>��A��I�z/������&���wKʺ�
�p��	ٕ�C�����?���8����R�郒��y�(�0��3;�����
�;�ť84��C�X��%�g���ϗ�	�Ϥ�6�B��j�ˉ" ?<�L2�A�>�g�⛈�q�?8`�n�1޵9��N S#5�ۭ�~Zvʲ�b��4X��U<��9�!s���;T�_�4y���x籠�>��'�)O\�xE}�~Eeý�����j���1N���hz�2���R�1z��Ä#��^,�0�M�M�$=rw�ɧ�`G)��0[bF��!԰���7�<g���0,$��vt��{��ϕS�㜱5{ng*YPR�_�o�~�����V0��__�̸���4j�v?c�{���;kW��c���+�*��6F�f��rꧼ`����[3����ь�o;�1U`�����1�#�%d��y)���TR��+�Q�$�����,#-�6�T���զ��{h�-�Pbw0��'tE��0�[�Ghi�X�s4w�Y�� >ۥ��Za�J$�8Q�z��Q,��� d"��)K��j�Cn[G�K28���iBڄ�+�~����D�I b��*HN�?�Kt/1�y��\�����l�7X7|��\��JdP(�^�|�]t	�L=9U�`�PS�����Tl�ݪGK�ɦ�=��Ƚ!8��V�k���-w^7(P��w�E(�����C�����}t
�"����(MC�Ѧ?�68!�!_�wJ:̝~%�8�z��|f� P*O�K"�{GM%dƺ�b�,�=*ړDk�0�#:��74�:T{��{xf�C_�3	Jj�$���z!a�*���$ZѾ�SD���a��nWʗSu�.��	�{|�o���'!/5އݯ���_*�i+� ��dg�\�ܸ�sǗ� �A���{0�å�Tf��ǽ�B���!�AǪ�,Gʨ��'�ä@���'�t�0��R�_Tv5P&h��j:��N4E�r��j��r�q�."X1��ޤ��%�a�t�%�g��D% �AТ�J����K�7���Gx��;0K��ê�{M�G�uܭfP��5�x�_]ǽA	u���i%D��a�S���;�M}[�,Th�
�����c�@��"���*��� O)a�y�5����=�4q����b��&�s%J�g��ѐ�,���UcQ_fUT"��/��B���V�Tc��[�]j�s�날�s�3e�+��*
h��_��8�c�<Di�OESA1�@'5+|W�
Kc� �5��ik�����J�T%����ׂ;�o�j8�U�:1���x-����O{�PZV���:�}�H�Y�T=�`Ju+w���-�'���,<�]��+�[�<*��U�LwĮ����_]h��
B��T?��t6<��Ԫ�v�%2>H.]\���]�qniA�+��B�:PT��#�������xvW�
��04Ŵ��m��z�C��%���Rj����b��?��6�<��I+��F.)�' XM�g?H��9Z���>�U��"���� b��m,�d�Q��J�+2��M�I��Z��'1�jw�A|a#�Jns$�h���m�4�u'�6�G��/�wo����k�bn��� �PuYHM"���P嬭+�|��+ҁ����I�> �\�9�J2� �K�қ��^T��2�:N�PQxK�.[��d�9jou�P[+�1�%A��R�4"�\�B'e�,tT뉫�7� _.�����d��K���^2B3Vɇ��N����.G1*��g(U��`��J�f�r2�k�zNC�B�=�=J�̪�Ec	�����a���8xt{��	�����j9�G�)蔎'3��z��̒��;K���w�7�ů��Bԛ�w�2�\�������I���_����h��J      �   �   x���A
�0E��)�/��ѐ��z���Mt���PP��W�i����Q����䌵,���������bw��0<M�F��k�S���)9_�!q�'�#�"�����B�j=R�e9�����X����-����h�ߴN$     