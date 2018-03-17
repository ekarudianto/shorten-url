require("expose-loader?$!expose-loader?jQuery!jquery");
require("bootstrap-sass/assets/javascripts/bootstrap.js");
$(() => {

    $('#form-url').submit(function() {
        var notifError = $('.summary.error');
        var notifErrorRequired = $('.summary.error.required');
        var notifSuccess = $('.summary.success');
        var shortenURLLink = $('#shortenURLLink');
        var url = $('#url');

        // Hide notifications
        notifError.hide();
        notifErrorRequired.hide();
        notifSuccess.hide();

        if (url.val() == '') {
            notifErrorRequired.show();
            return false;
        }

        $.ajax({
            dataType: 'json',
            type: 'POST',
            url: '/shorten',
            data: {
              url: url.val()
            },
            success: function(data) {
              var shortenURL = window.location.href + data.short_link;

              notifSuccess.show();
              shortenURLLink.attr('href', shortenURL);
              shortenURLLink.html(shortenURL);
            },
            error: function() {
              notifError.show();
            }
        });

        return false;
    });

});
