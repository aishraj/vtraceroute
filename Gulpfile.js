var gulp       = require('gulp'),

    browserify = require('gulp-browserify'),
    concat     = require('gulp-concat'),
    imagemin   = require('gulp-imagemin');
    del = require('del');

gulp.task('styles', function () {

    gulp.src(['public/assets/css/typography.css', 'public/assets/css/styles.css'])
        .pipe(concat('styles.css'))
        .pipe(gulp.dest('./public/build/'));

});

gulp.task('scripts', function () {

    gulp.src(['public/assets/js/app.js'])
        .pipe(browserify({
            debug: true,
            transform: [ 'babelify' ]
        }))
        .pipe(gulp.dest('public/build/'));

});

gulp.task('images', function () {

    gulp.src(['public/assets/img/**/*.png', 'public/assets/img/**/*.gif'])
        .pipe(imagemin())
        .pipe(gulp.dest('public/build/img/'));

});

gulp.task('dev', function () {

    gulp.run('build');

    gulp.watch('public/assets/js/**/*.js', [ 'scripts' ]);
    gulp.watch('public/assets/css/**/*.css', [ 'styles' ]);
    gulp.watch('public/assets/img/**/*', [ 'images' ]);

});

gulp.task('build', [ 'styles', 'scripts', 'images' ]);

gulp.task('clean', function(callback) {
  del(['./public/build'], callback)
});
