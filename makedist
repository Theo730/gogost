#!/bin/sh -ex

cur=$(pwd)
tmp=$(mktemp -d)
release=$1
[ -n "$release" ]

redo-ifchange streebog256
git clone . $tmp/gogost-$release
cd $tmp/gogost-$release
git checkout v$release
redo VERSION
go mod vendor
mkdir contrib
cp ~/work/redo/apenwarr/minimal/do contrib/do

cat > download.texi <<EOF
You can obtain releases source code prepared tarballs on
@url{http://www.gogost.cypherpunks.ru/}.
EOF

mkinfo() {
    ${MAKEINFO:-makeinfo} --plaintext \
        --set-customization-variable ASCII_PUNCTUATION=1 \
        -D "VERSION `cat VERSION`" $@
}

texi=$(mktemp)
cat > $texi <<EOF
\input texinfo
@documentencoding UTF-8
@settitle INSTALL
@include install.texi
@bye
EOF
mkinfo --output INSTALL $texi

cat > $texi <<EOF
\input texinfo
@documentencoding UTF-8
@settitle NEWS
@include news.texi
@bye
EOF
mkinfo --output NEWS $texi

cat > $texi <<EOF
\input texinfo
@documentencoding UTF-8
@settitle FAQ
@include faq.texi
@bye
EOF
mkinfo --output FAQ $texi

rm -rf .git
redo-cleanup full
rm -f \
    $texi \
    *.texi \
    .gitignore \
    clean.do \
    makedist.sh \
    style.css \
    TODO \
    VERSION.do \
    www.do

perl -i -npe "s/build/build -mod=vendor/" default.do
perl -i -npe "s/test/test -mod=vendor/" bench.do

find . -type d -exec chmod 755 {} +
find . -type f -exec chmod 644 {} +
chmod +x contrib/do

cd ..
tar cvf gogost-"$release".tar --uid=0 --gid=0 --numeric-owner gogost-"$release"
zstd -19 -v gogost-"$release".tar
tarball=gogost-"$release".tar.zst
gpg --detach-sign --sign --local-user 82343436696FC85A $tarball
gpg --enarmor < "$tarball".sig |
    sed "/^Comment:/d ; s/ARMORED FILE/SIGNATURE/" > "$tarball".asc
meta4-create -file "$tarball" -mtime "$tarball" -sig "$tarball".asc \
    http://www.gogost.cypherpunks.ru/"$tarball" \
    http://y.www.gogost.cypherpunks.ru/"$tarball" > "$tarball".meta4

size=$(( $(stat -f %z $tarball) / 1024 ))
hash=$(gpg --print-md SHA256 < $tarball)
hashsb=$($HOME/work/gogost/streebog256 < $tarball)
release_date=$(date "+%Y-%m-%d")

cat <<EOF
An entry for documentation:
@item @ref{Release $release, $release} @tab $release_date @tab $size KiB
@tab
    @url{$tarball.meta4, meta4}
    @url{$tarball, link}
    @url{$tarball.sig, sig}
@tab @code{$hash}
@tab @code{$hashsb}
EOF

cat <<EOF
Subject: [EN] GoGOST $release release announcement

I am pleased to announce GoGOST $release release availability!

GoGOST is free software pure Go GOST cryptographic functions library.
GOST is GOvernment STandard of Russian Federation (and Soviet Union).

------------------------ >8 ------------------------

The main improvements for that release are:


------------------------ >8 ------------------------

GoGOST'es home page is: http://www.gogost.cypherpunks.ru/

Source code and its signature for that version can be found here:

    http://www.gogost.cypherpunks.ru/gogost-${release}.tar.zst ($size KiB)
    http://www.gogost.cypherpunks.ru/gogost-${release}.tar.zst.sig

Streebog-256 hash: $hashsb
SHA256 hash: $hash
GPG key: CEBD 1282 2C46 9C02 A81A  0467 8234 3436 696F C85A
         GoGOST releases <gogost at cypherpunks dot ru>

Please send questions regarding the use of GoGOST, bug reports and patches
to mailing list: http://lists.cypherpunks.ru/gost.html
EOF

cat <<EOF
Subject: [RU] Состоялся релиз GoGOST $release

Я рад сообщить о выходе релиза GoGOST $release!

GoGOST это свободное программное обеспечение реализующее
криптографические функции ГОСТ на чистом Go. ГОСТ -- ГОсударственный
СТандарт Российской Федерации (а также Советского Союза).

------------------------ >8 ------------------------

Основные усовершенствования в этом релизе:


------------------------ >8 ------------------------

Домашняя страница GoGOST: http://www.gogost.cypherpunks.ru/

Исходный код и его подпись для этой версии могут быть найдены здесь:

    http://www.gogost.cypherpunks.ru/gogost-${release}.tar.zst ($size KiB)
    http://www.gogost.cypherpunks.ru/gogost-${release}.tar.zst.sig

Streebog-256 хэш: $hashsb
SHA256 хэш: $hash
GPG ключ: CEBD 1282 2C46 9C02 A81A  0467 8234 3436 696F C85A
          GoGOST releases <gogost at cypherpunks dot ru>

Пожалуйста, все вопросы касающиеся использования GoGOST, отчёты об
ошибках и патчи отправляйте в gost почтовую рассылку:
http://lists.cypherpunks.ru/gost.html
EOF

mv $tmp/$tarball $tmp/"$tarball".sig $tmp/"$tarball".meta4 $cur/gogost.html/
rm -fr $tmp
