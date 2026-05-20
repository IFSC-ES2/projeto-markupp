#!/bin/bash
# Script para instalar todas as dependências necessárias para compilar a apresentação

echo "Instalando dependências do LaTeX..."

sudo apt update

# Pacotes básicos do LaTeX
sudo apt install -y \
    texlive-latex-base \
    texlive-latex-extra \
    texlive-lang-portuguese \
    texlive-bibtex-extra \
    texlive-fonts-extra \
    texlive-science \
    biber

echo ""
echo "✓ Dependências instaladas!"
echo ""
echo "Agora você pode compilar a apresentação com:"
echo "  ./compilar.sh"
echo "ou"
echo "  make"

