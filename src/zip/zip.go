package zip

import (
    "archive/zip"
    "os"
    "io"
    "path/filepath"
    "lock/src/utils"
)

func ZipFiles(files []string) error {
    outFile, err := os.Create("out.zip")
    utils.Check(err)
    defer outFile.Close()

    zipWriter := zip.NewWriter(outFile)
    defer zipWriter.Close()

    for _, file := range files {
        zipfile, err := os.Open(file)
        utils.Check(err)
        defer zipfile.Close()

        // Get the file information
        info, err := zipfile.Stat()
        utils.Check(err)

        header, err := zip.FileInfoHeader(info)
        utils.Check(err)
        header.Method = zip.Deflate

        writer, err := zipWriter.CreateHeader(header)
        utils.Check(err)

        _, err = io.Copy(writer, zipfile)
        utils.Check(err)
    }
    return nil
}

// Unzip will decompress a zip archive, moving all files and folders 
// within the zip file (parameter 1) to an output directory (parameter 2).
func Unzip(src string, dest string) ([]string, error) {
    var filenames []string
    r, err := zip.OpenReader(src)
    utils.Check(err)
    defer r.Close()

    for _, f := range r.File {

        rc, err := f.Open()
        utils.Check(err)
        defer rc.Close()

        // Store filename/path for returning and using later on
        fpath := filepath.Join(dest, f.Name)
        filenames = append(filenames, fpath)

        if f.FileInfo().IsDir() {
            os.MkdirAll(fpath, os.ModePerm)
        } else {
            // Make File
            if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
                return filenames, err
            }

            outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
            utils.Check(err)

            _, err = io.Copy(outFile, rc)

            // Close the file without defer to close before next iteration of loop
            outFile.Close()
            utils.Check(err)
        }
    }
    return filenames, nil
}
