let 
  # commit revison of nixpkgs repository
  nixpkgsRev = "8f94e6566a75d02bc3910453f4c1d7bfcaeeb595";

  # a nix function to fetch a tar ball from github
  githubTarball = owner: repo: rev:
    builtins.fetchTarball { url = "https://github.com/${owner}/${repo}/archive/${rev}.tar.gz"; };

  # import set of packages from nixpkgs github repo at a particular revision with the config specified below
  pkgs = import (githubTarball "nixos" "nixpkgs" nixpkgsRev) { };
in
{
  inherit pkgs;
  shell = pkgs.mkShell {
    name = "go-lang";
    buildInputs = [
      pkgs.go
    ];
    shellHook = ''
      echo "Start learning go"
    '';
  };
}

