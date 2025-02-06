$(document).ready(function() {
    var sectionArray = [1, 2, 3, 4, 5];

    $.each(sectionArray, function(index, value) {
        var sectionElement = $('#' + 'section_' + value);

        if (sectionElement.length) { // ✅ ตรวจสอบว่า section มีอยู่จริงก่อนใช้ offset()
            $(document).scroll(function() {
                var offsetSection = sectionElement.offset().top - 94;
                var docScroll = $(document).scrollTop();
                var docScroll1 = docScroll + 1;

                if (docScroll1 >= offsetSection) {
                    $('.navbar-nav .nav-item .nav-link').removeClass('active');
                    $('.navbar-nav .nav-item .nav-link:link').addClass('inactive');  
                    $('.navbar-nav .nav-item .nav-link').eq(index).addClass('active');
                    $('.navbar-nav .nav-item .nav-link').eq(index).removeClass('inactive');
                }
            });

            $('.click-scroll').eq(index).click(function(e) {
                e.preventDefault();
                $('html, body').animate({ scrollTop: offsetSection }, 300);
            });
        } else {
            console.warn("Section " + value + " not found.");
        }
    });

    if ($(".scroll-to").length) {
        $(".scroll-to").click(function(event) {
            event.preventDefault();
            var target = $(this).attr("href");
            if ($(target).length) {
                $("html, body").animate({ scrollTop: $(target).offset().top }, 1000);
            }
        });
    } else {
        console.warn("No .scroll-to elements found.");
    }
});
