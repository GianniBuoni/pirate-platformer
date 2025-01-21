{
  description = "A very basic flake";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
  };

  outputs = {
    self,
    nixpkgs,
  }: let
    system = "x86_64-linux";
    pkgs = nixpkgs.legacyPackages.${system};
    gg = pkgs.writeShellScriptBin "gg" ''
      go run .
    '';
  in {
    devShells.${system}.default = pkgs.mkShell {
      packages =
        [
          gg
        ]
        ++ (with pkgs; [
          go
          tiled
        ]);

      shellHook = "tmux";
    };
  };
}
