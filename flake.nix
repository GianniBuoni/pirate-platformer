{
  description = "A raylib-go dev shell";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
  };

  outputs = {
    self,
    nixpkgs,
  }: let
    system = "x86_64-linux";
    pkgs = import nixpkgs {
      inherit system;
      config.allowUnfree = true;
    };
    ggg = pkgs.writeShellScriptBin "ggg" ''
      go run -tags x11 .
    '';
  in {
    devShells.${system}.default = pkgs.mkShell {
      packages =
        [
          ggg
        ]
        ++ (with pkgs; [
          go
          gopls
          libGL
          tiled
          aseprite

          # x11 dependencies
          xorg.libX11
          xorg.libX11.dev
          xorg.libXcursor
          xorg.libXi
          xorg.libXinerama
          xorg.libXrandr
        ]);

      buildInputs = with pkgs; [
        wayland
        glfw-wayland
        libxkbcommon
      ];
      shellHook = ''
        export LD_LIBRARY_PATH=${pkgs.wayland}/lib:$LD_LIBRARY_PATH
      '';
    };
  };
}
