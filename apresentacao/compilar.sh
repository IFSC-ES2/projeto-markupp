#!/bin/bash
# Script para compilar a apresentação LaTeX Beamer

cd "$(dirname "$0")"

# Nome do arquivo principal (sem extensão .tex)
ARQUIVO="apresentacao"

echo "Compilando apresentação LaTeX..."

# Primeira compilação
echo "Executando pdflatex (1/4)..."
pdflatex -interaction=nonstopmode "${ARQUIVO}.tex" > /dev/null 2>&1

# Processar bibliografia com biber
echo "Executando biber (2/4)..."
biber "${ARQUIVO}" > /dev/null 2>&1

# Segunda compilação (para resolver referências da bibliografia)
echo "Executando pdflatex (3/4)..."
pdflatex -interaction=nonstopmode "${ARQUIVO}.tex" > /dev/null 2>&1

# Terceira compilação (para garantir que tudo está resolvido)
echo "Executando pdflatex (4/4)..."
pdflatex -interaction=nonstopmode "${ARQUIVO}.tex" > /dev/null 2>&1

# Verificar se o PDF foi gerado
if [ -f "${ARQUIVO}.pdf" ]; then
    echo "✓ Compilação concluída com sucesso!"
    echo "✓ Arquivo gerado: ${ARQUIVO}.pdf"
    
    # Abrir o PDF (descomente a linha abaixo se desejar abrir automaticamente)
    # xdg-open "${ARQUIVO}.pdf" 2>/dev/null || evince "${ARQUIVO}.pdf" 2>/dev/null
else
    echo "✗ Erro: PDF não foi gerado. Verifique os logs acima."
    exit 1
fi

