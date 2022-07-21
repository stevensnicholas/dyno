/*
 * Sweet Alerts - Extra Components
 */

$(function() {
	"use strict";

	$('.btn-message').click(function() {
		swal("Here's a message!");
	});

	$('.btn-title-text').click(function() {
		swal("Here's a message!", "It's pretty, isn't it?")
	});

	$('.btn-timer').click(function() {
		swal({
			title: "Auto close alert!",
			text: "I will close in 2 seconds.",
			timer: 2000,
			showConfirmButton: false
		});
	});

	$('.btn-success').click(function() {
		swal("Good job!", "You clicked the button!", "success");
	});

	$('.btn-warning-confirm').click(function() {
		swal({
				title: "Are you sure?",
				text: "You will not be able to recover this imaginary file!",
				type: "warning",
				showCancelButton: true,
				confirmButtonColor: '#DD6B55',
				confirmButtonText: 'Yes, delete it!',
				closeOnConfirm: false
			},
			function() {
				swal("Deleted!", "Your imaginary file has been deleted!", "success");
			});
	});

	$('.btn-warning-cancel').click(function() {
		swal({
				title: "Are you sure?",
				text: "You will not be able to recover this imaginary file!",
				type: "warning",
				showCancelButton: true,
				confirmButtonColor: '#DD6B55',
				confirmButtonText: 'Yes, delete it!',
				cancelButtonText: "No, cancel plx!",
				closeOnConfirm: false,
				closeOnCancel: false
			},
			function(isConfirm) {
				if (isConfirm) {
					swal("Deleted!", "Your imaginary file has been deleted!", "success");
				} else {
					swal("Cancelled", "Your imaginary file is safe :)", "error");
				}
			});
	});

	$('.btn-custom-icon').click(function() {
		swal({
			title: "Sweet!",
			text: "Here's a custom image.",
			imageUrl: 'images/favicon/apple-touch-icon-152x152.png'
		});
	});

	$('.btn-message-html').click(function() {
		swal({
			title: "HTML <small>Title</small>!",
			text: 'A custom <span style="color:#F8BB86">html<span> message.',
			html: true
		});
	});

	$('.btn-input').click(function() {
		swal({
				title: "An input!",
				text: 'Write something interesting:',
				type: 'input',
				showCancelButton: true,
				closeOnConfirm: false,
				animation: "slide-from-top",
				inputPlaceholder: "Write something",
			},
			function(inputValue) {
				if (inputValue === false) return false;

				if (inputValue === "") {
					swal.showInputError("You need to write something!");
					return false;
				}

				swal("Nice!", 'You wrote: ' + inputValue, "success");

			});
	});

	$('.btn-theme').click(function() {
		swal({
			title: "Themes!",
			text: "Here's the Twitter theme for SweetAlert!",
			confirmButtonText: "Cool!",
			customClass: 'twitter'
		});
	});

	$('.btn-ajax').click(function() {
		swal({
			title: 'Ajax request example',
			text: 'Submit to run ajax request',
			type: 'info',
			showCancelButton: true,
			closeOnConfirm: false,
			showLoaderOnConfirm: true,
		}, function() {
			setTimeout(function() {
				swal('Ajax request finished!');
			}, 2000);
		});
	});

});