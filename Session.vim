let SessionLoad = 1
let s:so_save = &g:so | let s:siso_save = &g:siso | setg so=0 siso=0 | setl so=-1 siso=-1
let v:this_session=expand("<sfile>:p")
silent only
silent tabonly
cd ~/Projects/mech-commander
if expand('%') == '' && !&modified && line('$') <= 1 && getline(1) == ''
  let s:wipebuf = bufnr('%')
endif
let s:shortmess_save = &shortmess
if &shortmess =~ 'A'
  set shortmess=aoOA
else
  set shortmess=aoO
endif
badd +63 main.go
badd +12 commons/greeter_interface.go
badd +54 app.go
badd +1 plugin/greeter
badd +1 ~/Projects/mech-commander/go.mod
badd +1 ~/Projects/mech-commander/go.sum
badd +6 frontend/Session.vim
badd +976 ~/.asdf/installs/golang/1.19/go/src/syscall/zsyscall_darwin_arm64.go
badd +974 ~/.asdf/installs/golang/1.19.1/go/src/syscall/zsyscall_darwin_arm64.go
badd +388 ~/.asdf/installs/golang/1.19.1/go/src/os/file.go
badd +709 ~/.asdf/installs/golang/1.19.1/go/src/net/dial.go
badd +41 example/user_implementation.go
badd +80 ~/.asdf/installs/golang/1.19.1/go/src/os/tempfile.go
badd +122 ~/.asdf/installs/golang/1.19.1/go/src/net/net.go
badd +135 ~/.asdf/installs/golang/1.19.1/go/src/bufio/scan.go
badd +308 ~/.asdf/installs/golang/1.19.1/go/src/strings/strings.go
badd +39 ~/.asdf/installs/golang/1.19.1/packages/pkg/mod/github.com/wailsapp/wails/v2@v2.0.0/pkg/runtime/events.go
badd +2 frontend/src/App.svelte
argglobal
%argdel
edit app.go
argglobal
balt example/user_implementation.go
setlocal fdm=manual
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal fen
silent! normal! zE
let &fdl = &fdl
let s:l = 54 - ((24 * winheight(0) + 21) / 43)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 54
normal! 013|
tabnext 1
if exists('s:wipebuf') && len(win_findbuf(s:wipebuf)) == 0 && getbufvar(s:wipebuf, '&buftype') isnot# 'terminal'
  silent exe 'bwipe ' . s:wipebuf
endif
unlet! s:wipebuf
set winheight=1 winwidth=20
let &shortmess = s:shortmess_save
let s:sx = expand("<sfile>:p:r")."x.vim"
if filereadable(s:sx)
  exe "source " . fnameescape(s:sx)
endif
let &g:so = s:so_save | let &g:siso = s:siso_save
set hlsearch
let g:this_session = v:this_session
let g:this_obsession = v:this_session
doautoall SessionLoadPost
unlet SessionLoad
" vim: set ft=vim :
