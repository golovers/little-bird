var simplemde = new SimpleMDE({ 
    element: $("#postContent")[0],
    placeholder: "Enter content here...",
    renderingConfig: {
        singleLineBreaks: false,
        codeSyntaxHighlighting: true,
    },
});

function addArticle() {
    htmlValue = simplemde.markdown(simplemde.value())
    postData = {
        Title: $("#postTile").val(),
        Content: htmlValue,
        Markdown: simplemde.value(),
    }
    $.ajax({
        type: "POST",
        url: "/api/v1/articles",
        data: JSON.stringify(postData),
        contentType: "application/json; charset=utf-8",
        crossDomain: true,
        dataType: "json",
        success: function (data, status, jqXHR) {
            $(location).attr('href', '/articles/'+data.ID +'/details')
        },
        error: function (jqXHR, status) {
            alert('fail' + status.code);
        }
    });
}