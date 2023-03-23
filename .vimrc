set number
set hlsearch
syntax on

nnoremap <expr> <Right> (col('.') == col('$')-1 ? 'j0' : '<Right>')
nnoremap <expr> l (col('.') == col('$')-1 ? 'j0' : 'l')
inoremap <expr> <Right> (col('.') == col('$')-1 ? '<Esc>j0i' : '<Right>')
inoremap <expr> l (col('.') == col('$')-1 ? '<Esc>j0i' : 'l')

nnoremap <expr> <Left> (col('.') == 1 ? 'k$' : '<Left>')
nnoremap <expr> h (col('.') == 1 ? 'k$' : 'h')
inoremap <expr> <Left> (col('.') == 1 ? '<Esc>k$i' : '<Left>')
inoremap <expr> h (col('.') == 1 ? '<Esc>k$i' : 'h')
