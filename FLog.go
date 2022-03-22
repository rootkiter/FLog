/** ***********************************************
 * File Name : FLog.go
 * Author    : rootkiter
 * E-mail    : rootkiter@rootkiter.com
 * Created   : 2022-03-15 16:53:21 CST
************************************************* */

package FLog
import (
    "os"
    "fmt"
    "time"
    "path"
)

type FLog struct {
    directory string
    name      string
    nexttms   int64
    limittms  int64
    handle    *os.File
}

func (self *FLog) Init (
    directory string, name string, limittms int64,
) *FLog {
    self.directory = directory
    self.name      = name
    self.limittms  = limittms
    self.nexttms   = 0
    return self
}

func (self *FLog) WriteString (
    tms int64, logstr string,
) {
    if tms > self.nexttms || self.handle == nil {
        ftms := time.Unix(tms/1000,0).Format("20060102_150405_MST")
        fname := fmt.Sprintf("%s_%s.json", self.name, ftms)
        fpath := path.Join(self.directory, fname)
        handle, err := os.OpenFile(
            fpath,
            os.O_CREATE|os.O_WRONLY|os.O_APPEND,
            0644,
        )
        if err != nil {
            self.handle = nil
        } else {
            self.Close()
            self.handle = handle
        }
        self.nexttms = tms + self.limittms
    }
    if self.handle != nil {
        self.handle.WriteString(logstr)
    }
}

func (self *FLog) WriteLogString ( logstr string ) {
    tms := time.Now().Unix()*1000
    logstring := fmt.Sprintf("%d\t%s\n", tms/1000, logstr)
    self.WriteString(tms, logstring)
}

func (self *FLog) Close () {
    if self.handle != nil {
        self.handle.Close()
        self.handle = nil
    }
}
