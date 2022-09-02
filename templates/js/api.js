var urlRoot = "http://127.0.0.1:8080";
var urlBackToRootPath = `${urlRoot}/BackToRootPath`;
var urlBackToPrevPath = `${urlRoot}/BackToPrevPath`;
var urlJoinNextPath = `${urlRoot}/JoinNextPath`;
var urlDeleteFile = `${urlRoot}/DeleteFile`;
var urlReloadCurPath = `${urlRoot}/ReloadCurPath`;
var urlUploadFile = `${urlRoot}/UploadFile`;
var urlDownloadFile = `${urlRoot}/DownloadFile`;

function ApiDeleteFile(filefullpath, tableIns) {
  axios
    .delete(urlDeleteFile, {
      data: {
        delete_file_path: filefullpath,
      },
    })
    .then(function (response) {
      if (response.data.code == 0) {
        tableIns.reload({
          url: urlReloadCurPath,
        });
      } else {
        layer.msg(JSON.stringify(response.data));
      }
    })
    .catch(function (error) {
      console.log(error);
    });
}

function ApiDownloadFile(data) {
  axios({
    method: "post",
    url: urlDownloadFile,
    data: { download_file_path: data.filefullpath },
    responseType: "blob",
  }).then((res) => {
    console.log(res);
    // 处理返回的文件流
    const content = res.data;
    const blob = new Blob([content]);
    const fileName = data.filename;
    if ("download" in document.createElement("a")) {
      // 非IE下载
      const elink = document.createElement("a");
      elink.download = fileName;
      elink.style.display = "none";
      elink.href = URL.createObjectURL(blob);
      document.body.appendChild(elink);
      elink.click();
      URL.revokeObjectURL(elink.href);
      document.body.removeChild(elink);
    } else {
      // IE10+下载
      navigator.msSaveBlob(blob, fileName);
    }
  });
}
