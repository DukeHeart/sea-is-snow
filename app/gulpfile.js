var gulp = require('gulp');

var watch = require('gulp-watch');
var livereload = require('gulp-livereload');
gulp.task('watch', function() {
  livereload.listen();
  gulp.watch('../www/**.*').on('change', livereload.changed);
});

var shell = require('gulp-shell')
gulp.task('appserver', shell.task([
  'dev_appserver.py app.yaml'
]));

var shell = require('gulp-shell')
gulp.task('update', shell.task([
  'vulcanize -o www/build.html ../www/index.html --inline --strip',
  'appcfg.py --oauth2 update app.yaml'
]));

var shell = require('gulp-shell')
gulp.task('dispatch', shell.task([
  'appcfg.py --oauth2 update_dispatch dispatch.yaml'
]));

var shell = require('gulp-shell')
gulp.task('index', shell.task([
  'appcfg.py --oauth2 update_indexes index.yaml'
]));

var shell = require('gulp-shell')
gulp.task('vacuum', shell.task([
  'appcfg.py --oauth2 vacuum_indexes index.yaml'
]));

var shell = require('gulp-shell')
gulp.task('rollback', shell.task([
  'appcfg.py --oauth2 rollback app.yaml'
]));

gulp.task('build', shell.task([
  'vulcanize -o ../www/build.html ../www/index.html --inline --strip'
]));

gulp.task('test', shell.task([
  'vulcanize -o ../www/testB.html ../www/test.html --inline --strip'
]));

gulp.task('default', ['watch', 'appserver']);

/*
var vulcanize = require('gulp-vulcanize');
gulp.task('build', function() {
  return gulp.src('../www/index.html')
    .pipe(vulcanize({dest: './', inline: true, strip: true }))
    .pipe(gulp.dest('./'));
 });
*/

/*
var webserver = require('gulp-webserver');
gulp.task('webserver', function() {
  gulp.src('../www')
    .pipe(webserver({
      livereload: true,
      directoryListing: true,
      open: true,
      fallback: 'index.html'
    }));
});
*/