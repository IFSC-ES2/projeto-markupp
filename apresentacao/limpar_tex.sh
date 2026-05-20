#!/bin/bash
# Script para limpar arquivos temporários do LaTeX

cd "$(dirname "$0")"

# Remove arquivos temporários
rm -f *.aux *.bbl *.bcf *.blg *.fls *.log *.nav *.out *.run.xml *.snm *.synctex.gz *.toc *.vrb 2>/dev/null
