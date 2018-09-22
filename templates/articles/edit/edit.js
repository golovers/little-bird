function updateArticle() {
    htmlValue = simplemde.markdown(simplemde.value())
    article = {
        ID: $("#id").val(),
        Title: $("#title").val(),
        Content: htmlValue,
        Markdown: simplemde.value(),
    }
    $.ajax({
        type: "PUT",
        url: "/api/v1/articles/" + article.ID,
        data: JSON.stringify(article),
        contentType: "application/json; charset=utf-8",
        crossDomain: true,
        dataType: "json",
        success: function (data, status, jqXHR) {
            $(location).attr('href', '/articles/'+data.ID + "/details")
        },
        error: function (jqXHR, status) {
            alert('fail' + status.code);
        }
    });
}