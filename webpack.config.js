var webpack = require('webpack');
var path = require('path');
var HtmlWebpackPlugin = require('html-webpack-plugin')
var ExtractTextPlugin = require('extract-text-webpack-plugin');

module.exports = {
  entry: {
    'app': './static/app/boot.ts',
    'vendor': [
        './static/app/vendor.ts'
    ]
  },
  output: {
    path: "./webapp",
    filename: "[name].js"
  },
  debug: true,
  devtool: 'source-map',
  plugins: [
    new webpack.optimize.CommonsChunkPlugin('vendor', 'vendor.js'),
    new HtmlWebpackPlugin({
        template: "./static/index.html",
        inject: "body"
    }),
    new webpack.ProvidePlugin({
        $: "jquery",
        jQuery: "jquery",
        "window.jQuery": "jquery"
    }),
    new ExtractTextPlugin('[name].css'),
  ],

  resolve: {
    extensions: ['', '.ts', '.js']
  },

  module: {
    loaders: [
      { test: /\.ts$/,                loader: 'ts-loader' },
      { test: /\.(png|gif|jpg|svg)$/, loader: 'file?name=imgs/[name].[ext]?[hash]' },
      { test: /\.(eot|ttf|woff2?)$/,  loader: 'file?name=fonts/[name].[ext]?[hash]' },
      { test: /\.css$/,               loader: ExtractTextPlugin.extract('css?minimize') }, 
      { test: /\.less$/,              loader: ExtractTextPlugin.extract('css?minimize!less') }
    ],
    noParse: [ path.join(__dirname, 'node_modules', 'angular2', 'bundles') ]
  },

  devServer: {
    historyApiFallback: false
  }
};
