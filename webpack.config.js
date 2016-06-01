const path = require('path');
const Webpack = require('webpack');

module.exports = {
  entry: ['./public/js/main.js'],
  output: {
    path: `${__dirname}/dist`,
    publicPath: 'http://localhost:9000/public/',
    filename: './js/[name].js',
  },
  module: {
    preLoaders: [
      {
        test: /\.js$/,
        loader: 'eslint-loader',
        exclude: /node_modules/,
      },
    ],
    loaders: [
      {
        test: /\.(ttf|otf|eot|svg|png|jpg|gif)$/,
        loader: 'file-loader?name=img/[name].[ext]',
      },
      {
        test: /\.woff(2)?(\?v=[0-9]\.[0-9]\.[0-9]|)?$/,
        loader: 'url-loader?limit=10000&mimetype=application/font-woff',
      },
      {
        test: /\.js$/,
        exclude: /node_modules/,
        loader: 'babel',
        query: {
          presets: ['es2015', 'react'],
        },
      },
      {
        test: /\.css$/,
        loader: 'style-loader!css-loader',
      },
      {
        test: /\.sass$/,
        loader: 'style-loader!css-loader!sass-loader',
      },
    ],
  },
  resolve: {
    root: [
      path.resolve(__dirname, './public/js/'),
    ],
  },
  sassLoader: {
    includePaths: [path.resolve(__dirname, './public/sass')],
  },
  plugins: [
    new Webpack.ProvidePlugin({
      $: 'jquery',
      jQuery: 'jquery',
    }),
  ],
};
