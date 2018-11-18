package main

/*
#cgo CFLAGS: -fno-keep-inline-dllexport -O2 -Wall -W -Wextra -DUNICODE -D_UNICODE -DQT_NEEDS_QMAIN -DQT_NO_DEBUG -DQT_GUI_LIB -DQT_CORE_LIB
#cgo CXXFLAGS: -O2 -std=gnu++11 -Wall -W -fexceptions -mthreads -DUNICODE -D_UNICODE -DQT_NEEDS_QMAIN -DQT_NO_DEBUG -DQT_GUI_LIB -DQT_CORE_LIB
#cgo CXXFLAGS: -I../../src -I. -IE:/Qt/5.9.4/mingw53_32/include -IE:/Qt/5.9.4/mingw53_32/include/QtGui -IE:/Qt/5.9.4/mingw53_32/include/QtANGLE -IE:/Qt/5.9.4/mingw53_32/include/QtCore -Irelease -IE:/Qt/5.9.4/mingw53_32/mkspecs/win32-g++
#cgo LDFLAGS:        -Wl,-s -Wl,-subsystem,windows -mthreads
#cgo LDFLAGS:        -lmingw32 -LE:/Qt/5.9.4/mingw53_32/lib E:/Qt/5.9.4/mingw53_32/lib/libqtmain.a -LC:/utils/my_sql/my_sql/lib -LC:/utils/postgresql/pgsql/lib -lshell32 E:/Qt/5.9.4/mingw53_32/lib/libQt5Gui.a E:/Qt/5.9.4/mingw53_32/lib/libQt5Core.a
#cgo LDFLAGS: -Wl,--allow-multiple-definition
#cgo CFLAGS: -Wno-unused-parameter -Wno-unused-variable -Wno-return-type
#cgo CXXFLAGS: -Wno-unused-parameter -Wno-unused-variable -Wno-return-type
*/
import "C"
