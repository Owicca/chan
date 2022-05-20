package media

//_, err := mp4.ReadBoxStructure(f, func(h *mp4.ReadHandle) (any, error) {
//	log.Println("depth", len(h.Path))
//	log.Println("type", h.BoxInfo.Type.String())
//	log.Println("size", h.BoxInfo.Size)

//	return nil, nil
//})
//if err != nil {
//	logs.LogErr(op, errors.Errorf("err while parsing %s (%s)", mime, err))
//}

//filePath := f.(*os.File).Name()
//cmdStr := fmt.Sprintf(`ffprobe -v error -select_streams v:0 -show_entries stream=width,height -of csv=s=x:p=0 "%s"`, filePath)
//cmd := exec.Command(cmdStr)
//if err := cmd.Start(); err != nil {
//	log.Println("err in probing file", err)
//	return x, y
//}
//res, err := cmd.Output()
//if err != nil {
//	log.Println("err getting output of a file probing", err)
//	return x, y
//}
//log.Println("output", string(res))
