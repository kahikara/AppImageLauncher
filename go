mkdir -p ~/pkg/appimagelauncher-git-local
cd ~/pkg/appimagelauncher-git-local || exit 1

cat > PKGBUILD <<'EOF'
pkgname=appimagelauncher-git-local
_pkgname=AppImageLauncher
pkgver=0
pkgrel=1
pkgdesc="A Helper application for running and integrating AppImages"
arch=(x86_64)
url="https://github.com/TheAssassin/AppImageLauncher"
license=(MIT)

depends=(
  qt5-base
  qt5-declarative
  fuse2
  glibc
  gcc-libs
  glib2
  zstd
  cairo
  librsvg
  xz
  libarchive
  zlib
)

makedepends=(
  git
  cmake
  boost
  qt5-tools
  libxpm
  lib32-glibc
  lib32-gcc-libs
  xxd
  patchelf
  argagg
  nlohmann-json
)

optdepends=(
  'qt5-wayland: Qt 5 Wayland platform plugin'
)

provides=(appimagelauncher)
conflicts=(appimagelauncher appimagelauncher-git)

source=("AppImageLauncher::git+file:///home/ming/github/AppImageLauncher")
sha256sums=('SKIP')

pkgver() {
  cd "$srcdir/AppImageLauncher"
  printf "r%s.%s" "$(git rev-list --count HEAD)" "$(git rev-parse --short HEAD)"
}

build() {
  export CFLAGS+=" -w"
  export CXXFLAGS+=" -w"
  export CMAKE_POLICY_VERSION_MINIMUM=3.5

  cmake -B build -S "$srcdir/AppImageLauncher" -Wno-dev \
    -DCMAKE_BUILD_TYPE=Release \
    -DCMAKE_INSTALL_PREFIX=/usr \
    -DBUILD_TESTING=OFF \
    -DFETCHCONTENT_QUIET=OFF \
    -DCPR_FORCE_USE_SYSTEM_CURL=ON

  cmake --build build -j"$(nproc)"
}

package() {
  DESTDIR="$pkgdir" cmake --install build
  install -Dm644 "$srcdir/AppImageLauncher/LICENSE.txt" "$pkgdir/usr/share/licenses/$pkgname/LICENSE"
}
EOF

makepkg -si
