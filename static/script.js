$(document).ready(function(){
    $("#download-button").click(function(){
      console.log("Download button was clicked.");
      // $.get("/download", function(data, status){
      //   console.log("get download")
      // });
      $.ajax({
        url: '/download',
        method: 'GET',
        xhrFields: {
            responseType: 'blob'
        },
        success: function (data) {
            window.history.replaceState('data to be passed', 'Title of the page', '/download');
            var a = document.createElement('a');
            var url = window.URL.createObjectURL(data);
            a.href = url;
            a.download = 'download.txt';
            document.body.append(a);
            a.click();
            a.remove();
            window.URL.revokeObjectURL(url);
        }
      });
    });

    $("#back-button").click(function() {
      console.log("Back button was clicked.");
      window.history.back();
      //window.history.replaceState('data to be passed', 'Title of the page', '/');
      window.location.replace("/");
      console.log($("#text"))
      
    });
    
  });

