;(function(){

/**
 * Require the given path.
 *
 * @param {String} path
 * @return {Object} exports
 * @api public
 */

function require(path, parent, orig) {
  var resolved = require.resolve(path);

  // lookup failed
  if (null == resolved) {
    orig = orig || path;
    parent = parent || 'root';
    var err = new Error('Failed to require "' + orig + '" from "' + parent + '"');
    err.path = orig;
    err.parent = parent;
    err.require = true;
    throw err;
  }

  var module = require.modules[resolved];

  // perform real require()
  // by invoking the module's
  // registered function
  if (!module.exports) {
    module.exports = {};
    module.client = module.component = true;
    module.call(this, module.exports, require.relative(resolved), module);
  }

  return module.exports;
}

/**
 * Registered modules.
 */

require.modules = {};

/**
 * Registered aliases.
 */

require.aliases = {};

/**
 * Resolve `path`.
 *
 * Lookup:
 *
 *   - PATH/index.js
 *   - PATH.js
 *   - PATH
 *
 * @param {String} path
 * @return {String} path or null
 * @api private
 */

require.resolve = function(path) {
  if (path.charAt(0) === '/') path = path.slice(1);

  var paths = [
    path,
    path + '.js',
    path + '.json',
    path + '/index.js',
    path + '/index.json'
  ];

  for (var i = 0; i < paths.length; i++) {
    var path = paths[i];
    if (require.modules.hasOwnProperty(path)) return path;
    if (require.aliases.hasOwnProperty(path)) return require.aliases[path];
  }
};

/**
 * Normalize `path` relative to the current path.
 *
 * @param {String} curr
 * @param {String} path
 * @return {String}
 * @api private
 */

require.normalize = function(curr, path) {
  var segs = [];

  if ('.' != path.charAt(0)) return path;

  curr = curr.split('/');
  path = path.split('/');

  for (var i = 0; i < path.length; ++i) {
    if ('..' == path[i]) {
      curr.pop();
    } else if ('.' != path[i] && '' != path[i]) {
      segs.push(path[i]);
    }
  }

  return curr.concat(segs).join('/');
};

/**
 * Register module at `path` with callback `definition`.
 *
 * @param {String} path
 * @param {Function} definition
 * @api private
 */

require.register = function(path, definition) {
  require.modules[path] = definition;
};

/**
 * Alias a module definition.
 *
 * @param {String} from
 * @param {String} to
 * @api private
 */

require.alias = function(from, to) {
  if (!require.modules.hasOwnProperty(from)) {
    throw new Error('Failed to alias "' + from + '", it does not exist');
  }
  require.aliases[to] = from;
};

/**
 * Return a require function relative to the `parent` path.
 *
 * @param {String} parent
 * @return {Function}
 * @api private
 */

require.relative = function(parent) {
  var p = require.normalize(parent, '..');

  /**
   * lastIndexOf helper.
   */

  function lastIndexOf(arr, obj) {
    var i = arr.length;
    while (i--) {
      if (arr[i] === obj) return i;
    }
    return -1;
  }

  /**
   * The relative require() itself.
   */

  function localRequire(path) {
    var resolved = localRequire.resolve(path);
    return require(resolved, parent, path);
  }

  /**
   * Resolve relative to the parent.
   */

  localRequire.resolve = function(path) {
    var c = path.charAt(0);
    if ('/' == c) return path.slice(1);
    if ('.' == c) return require.normalize(p, path);

    // resolve deps by returning
    // the dep in the nearest "deps"
    // directory
    var segs = parent.split('/');
    var i = lastIndexOf(segs, 'deps') + 1;
    if (!i) i = 0;
    path = segs.slice(0, i + 1).join('/') + '/deps/' + path;
    return path;
  };

  /**
   * Check if module is defined at `path`.
   */

  localRequire.exists = function(path) {
    return require.modules.hasOwnProperty(localRequire.resolve(path));
  };

  return localRequire;
};
require.register("avetisk-defaults/index.js", function(exports, require, module){
'use strict';

/**
 * Merge default values.
 *
 * @param {Object} dest
 * @param {Object} defaults
 * @return {Object}
 * @api public
 */
var defaults = function (dest, src, recursive) {
  for (var prop in src) {
    if (recursive && dest[prop] instanceof Object && src[prop] instanceof Object) {
      dest[prop] = defaults(dest[prop], src[prop], true);
    } else if (! (prop in dest)) {
      dest[prop] = src[prop];
    }
  }

  return dest;
};

/**
 * Expose `defaults`.
 */
module.exports = defaults;

});
require.register("component-bind/index.js", function(exports, require, module){

/**
 * Slice reference.
 */

var slice = [].slice;

/**
 * Bind `obj` to `fn`.
 *
 * @param {Object} obj
 * @param {Function|String} fn or string
 * @return {Function}
 * @api public
 */

module.exports = function(obj, fn){
  if ('string' == typeof fn) fn = obj[fn];
  if ('function' != typeof fn) throw new Error('bind() requires a function');
  var args = [].slice.call(arguments, 2);
  return function(){
    return fn.apply(obj, args.concat(slice.call(arguments)));
  }
};

});
require.register("component-clone/index.js", function(exports, require, module){

/**
 * Module dependencies.
 */

var type;

try {
  type = require('type');
} catch(e){
  type = require('type-component');
}

/**
 * Module exports.
 */

module.exports = clone;

/**
 * Clones objects.
 *
 * @param {Mixed} any object
 * @api public
 */

function clone(obj){
  switch (type(obj)) {
    case 'object':
      var copy = {};
      for (var key in obj) {
        if (obj.hasOwnProperty(key)) {
          copy[key] = clone(obj[key]);
        }
      }
      return copy;

    case 'array':
      var copy = new Array(obj.length);
      for (var i = 0, l = obj.length; i < l; i++) {
        copy[i] = clone(obj[i]);
      }
      return copy;

    case 'regexp':
      // from millermedeiros/amd-utils - MIT
      var flags = '';
      flags += obj.multiline ? 'm' : '';
      flags += obj.global ? 'g' : '';
      flags += obj.ignoreCase ? 'i' : '';
      return new RegExp(obj.source, flags);

    case 'date':
      return new Date(obj.getTime());

    default: // string, number, boolean, …
      return obj;
  }
}

});
require.register("component-cookie/index.js", function(exports, require, module){
/**
 * Encode.
 */

var encode = encodeURIComponent;

/**
 * Decode.
 */

var decode = decodeURIComponent;

/**
 * Set or get cookie `name` with `value` and `options` object.
 *
 * @param {String} name
 * @param {String} value
 * @param {Object} options
 * @return {Mixed}
 * @api public
 */

module.exports = function(name, value, options){
  switch (arguments.length) {
    case 3:
    case 2:
      return set(name, value, options);
    case 1:
      return get(name);
    default:
      return all();
  }
};

/**
 * Set cookie `name` to `value`.
 *
 * @param {String} name
 * @param {String} value
 * @param {Object} options
 * @api private
 */

function set(name, value, options) {
  options = options || {};
  var str = encode(name) + '=' + encode(value);

  if (null == value) options.maxage = -1;

  if (options.maxage) {
    options.expires = new Date(+new Date + options.maxage);
  }

  if (options.path) str += '; path=' + options.path;
  if (options.domain) str += '; domain=' + options.domain;
  if (options.expires) str += '; expires=' + options.expires.toGMTString();
  if (options.secure) str += '; secure';

  document.cookie = str;
}

/**
 * Return all cookies.
 *
 * @return {Object}
 * @api private
 */

function all() {
  return parse(document.cookie);
}

/**
 * Get cookie `name`.
 *
 * @param {String} name
 * @return {String}
 * @api private
 */

function get(name) {
  return all()[name];
}

/**
 * Parse cookie `str`.
 *
 * @param {String} str
 * @return {Object}
 * @api private
 */

function parse(str) {
  var obj = {};
  var pairs = str.split(/ *; */);
  var pair;
  if ('' == pairs[0]) return obj;
  for (var i = 0; i < pairs.length; ++i) {
    pair = pairs[i].split('=');
    obj[decode(pair[0])] = decode(pair[1]);
  }
  return obj;
}

});
require.register("component-props/index.js", function(exports, require, module){
/**
 * Global Names
 */

var globals = /\b(Array|Date|Object|Math|JSON)\b/g;

/**
 * Return immediate identifiers parsed from `str`.
 *
 * @param {String} str
 * @param {String|Function} map function or prefix
 * @return {Array}
 * @api public
 */

module.exports = function(str, fn){
  var p = unique(props(str));
  if (fn && 'string' == typeof fn) fn = prefixed(fn);
  if (fn) return map(str, p, fn);
  return p;
};

/**
 * Return immediate identifiers in `str`.
 *
 * @param {String} str
 * @return {Array}
 * @api private
 */

function props(str) {
  return str
    .replace(/\.\w+|\w+ *\(|"[^"]*"|'[^']*'|\/([^/]+)\//g, '')
    .replace(globals, '')
    .match(/[a-zA-Z_]\w*/g)
    || [];
}

/**
 * Return `str` with `props` mapped with `fn`.
 *
 * @param {String} str
 * @param {Array} props
 * @param {Function} fn
 * @return {String}
 * @api private
 */

function map(str, props, fn) {
  var re = /\.\w+|\w+ *\(|"[^"]*"|'[^']*'|\/([^/]+)\/|[a-zA-Z_]\w*/g;
  return str.replace(re, function(_){
    if ('(' == _[_.length - 1]) return fn(_);
    if (!~props.indexOf(_)) return _;
    return fn(_);
  });
}

/**
 * Return unique array.
 *
 * @param {Array} arr
 * @return {Array}
 * @api private
 */

function unique(arr) {
  var ret = [];

  for (var i = 0; i < arr.length; i++) {
    if (~ret.indexOf(arr[i])) continue;
    ret.push(arr[i]);
  }

  return ret;
}

/**
 * Map with prefix `str`.
 */

function prefixed(str) {
  return function(_){
    return str + _;
  };
}

});
require.register("component-to-function/index.js", function(exports, require, module){
/**
 * Module Dependencies
 */

var expr = require('props');

/**
 * Expose `toFunction()`.
 */

module.exports = toFunction;

/**
 * Convert `obj` to a `Function`.
 *
 * @param {Mixed} obj
 * @return {Function}
 * @api private
 */

function toFunction(obj) {
  switch ({}.toString.call(obj)) {
    case '[object Object]':
      return objectToFunction(obj);
    case '[object Function]':
      return obj;
    case '[object String]':
      return stringToFunction(obj);
    case '[object RegExp]':
      return regexpToFunction(obj);
    default:
      return defaultToFunction(obj);
  }
}

/**
 * Default to strict equality.
 *
 * @param {Mixed} val
 * @return {Function}
 * @api private
 */

function defaultToFunction(val) {
  return function(obj){
    return val === obj;
  }
}

/**
 * Convert `re` to a function.
 *
 * @param {RegExp} re
 * @return {Function}
 * @api private
 */

function regexpToFunction(re) {
  return function(obj){
    return re.test(obj);
  }
}

/**
 * Convert property `str` to a function.
 *
 * @param {String} str
 * @return {Function}
 * @api private
 */

function stringToFunction(str) {
  // immediate such as "> 20"
  if (/^ *\W+/.test(str)) return new Function('_', 'return _ ' + str);

  // properties such as "name.first" or "age > 18" or "age > 18 && age < 36"
  return new Function('_', 'return ' + get(str));
}

/**
 * Convert `object` to a function.
 *
 * @param {Object} object
 * @return {Function}
 * @api private
 */

function objectToFunction(obj) {
  var match = {}
  for (var key in obj) {
    match[key] = typeof obj[key] === 'string'
      ? defaultToFunction(obj[key])
      : toFunction(obj[key])
  }
  return function(val){
    if (typeof val !== 'object') return false;
    for (var key in match) {
      if (!(key in val)) return false;
      if (!match[key](val[key])) return false;
    }
    return true;
  }
}

/**
 * Built the getter function. Supports getter style functions
 *
 * @param {String} str
 * @return {String}
 * @api private
 */

function get(str) {
  var props = expr(str);
  if (!props.length) return '_.' + str;

  var val;
  for(var i = 0, prop; prop = props[i]; i++) {
    val = '_.' + prop;
    val = "('function' == typeof " + val + " ? " + val + "() : " + val + ")";
    str = str.replace(new RegExp(prop, 'g'), val);
  }

  return str;
}

});
require.register("component-each/index.js", function(exports, require, module){

/**
 * Module dependencies.
 */

var toFunction = require('to-function');
var type;

try {
  type = require('type-component');
} catch (e) {
  type = require('type');
}

/**
 * HOP reference.
 */

var has = Object.prototype.hasOwnProperty;

/**
 * Iterate the given `obj` and invoke `fn(val, i)`.
 *
 * @param {String|Array|Object} obj
 * @param {Function} fn
 * @api public
 */

module.exports = function(obj, fn){
  fn = toFunction(fn);
  switch (type(obj)) {
    case 'array':
      return array(obj, fn);
    case 'object':
      if ('number' == typeof obj.length) return array(obj, fn);
      return object(obj, fn);
    case 'string':
      return string(obj, fn);
  }
};

/**
 * Iterate string chars.
 *
 * @param {String} obj
 * @param {Function} fn
 * @api private
 */

function string(obj, fn) {
  for (var i = 0; i < obj.length; ++i) {
    fn(obj.charAt(i), i);
  }
}

/**
 * Iterate object keys.
 *
 * @param {Object} obj
 * @param {Function} fn
 * @api private
 */

function object(obj, fn) {
  for (var key in obj) {
    if (has.call(obj, key)) {
      fn(key, obj[key]);
    }
  }
}

/**
 * Iterate array-ish.
 *
 * @param {Array|Object} obj
 * @param {Function} fn
 * @api private
 */

function array(obj, fn) {
  for (var i = 0; i < obj.length; ++i) {
    fn(obj[i], i);
  }
}

});
require.register("component-trim/index.js", function(exports, require, module){

exports = module.exports = trim;

function trim(str){
  if (str.trim) return str.trim();
  return str.replace(/^\s*|\s*$/g, '');
}

exports.left = function(str){
  if (str.trimLeft) return str.trimLeft();
  return str.replace(/^\s*/, '');
};

exports.right = function(str){
  if (str.trimRight) return str.trimRight();
  return str.replace(/\s*$/, '');
};

});
require.register("component-querystring/index.js", function(exports, require, module){

/**
 * Module dependencies.
 */

var trim = require('trim');

/**
 * Parse the given query `str`.
 *
 * @param {String} str
 * @return {Object}
 * @api public
 */

exports.parse = function(str){
  if ('string' != typeof str) return {};

  str = trim(str);
  if ('' == str) return {};

  var obj = {};
  var pairs = str.split('&');
  for (var i = 0; i < pairs.length; i++) {
    var parts = pairs[i].split('=');
    obj[parts[0]] = null == parts[1]
      ? ''
      : decodeURIComponent(parts[1]);
  }

  return obj;
};

/**
 * Stringify the given `obj`.
 *
 * @param {Object} obj
 * @return {String}
 * @api public
 */

exports.stringify = function(obj){
  if (!obj) return '';
  var pairs = [];
  for (var key in obj) {
    pairs.push(encodeURIComponent(key) + '=' + encodeURIComponent(obj[key]));
  }
  return pairs.join('&');
};

});
require.register("component-type/index.js", function(exports, require, module){

/**
 * toString ref.
 */

var toString = Object.prototype.toString;

/**
 * Return the type of `val`.
 *
 * @param {Mixed} val
 * @return {String}
 * @api public
 */

module.exports = function(val){
  switch (toString.call(val)) {
    case '[object Function]': return 'function';
    case '[object Date]': return 'date';
    case '[object RegExp]': return 'regexp';
    case '[object Arguments]': return 'arguments';
    case '[object Array]': return 'array';
    case '[object String]': return 'string';
  }

  if (val === null) return 'null';
  if (val === undefined) return 'undefined';
  if (val && val.nodeType === 1) return 'element';
  if (val === Object(val)) return 'object';

  return typeof val;
};

});
require.register("ianstormtaylor-is-empty/index.js", function(exports, require, module){

/**
 * Expose `isEmpty`.
 */

module.exports = isEmpty;


/**
 * Has.
 */

var has = Object.prototype.hasOwnProperty;


/**
 * Test whether a value is "empty".
 *
 * @param {Mixed} val
 * @return {Boolean}
 */

function isEmpty (val) {
  if (null == val) return true;
  if ('number' == typeof val) return 0 === val;
  if (undefined !== val.length) return 0 === val.length;
  for (var key in val) if (has.call(val, key)) return false;
  return true;
}
});
require.register("segmentio-extend/index.js", function(exports, require, module){

module.exports = function extend (object) {
    // Takes an unlimited number of extenders.
    var args = Array.prototype.slice.call(arguments, 1);

    // For each extender, copy their properties on our object.
    for (var i = 0, source; source = args[i]; i++) {
        if (!source) continue;
        for (var property in source) {
            object[property] = source[property];
        }
    }

    return object;
};
});
require.register("component-url/index.js", function(exports, require, module){

/**
 * Parse the given `url`.
 *
 * @param {String} str
 * @return {Object}
 * @api public
 */

exports.parse = function(url){
  var a = document.createElement('a');
  a.href = url;
  return {
    href: a.href,
    host: a.host || location.host,
    port: ('0' === a.port || '' === a.port) ? port(a.protocol) : a.port,
    hash: a.hash,
    hostname: a.hostname || location.hostname,
    pathname: a.pathname.charAt(0) != '/' ? '/' + a.pathname : a.pathname,
    protocol: !a.protocol || ':' == a.protocol ? location.protocol : a.protocol,
    search: a.search,
    query: a.search.slice(1)
  };
};

/**
 * Check if `url` is absolute.
 *
 * @param {String} url
 * @return {Boolean}
 * @api public
 */

exports.isAbsolute = function(url){
  return 0 == url.indexOf('//') || !!~url.indexOf('://');
};

/**
 * Check if `url` is relative.
 *
 * @param {String} url
 * @return {Boolean}
 * @api public
 */

exports.isRelative = function(url){
  return !exports.isAbsolute(url);
};

/**
 * Check if `url` is cross domain.
 *
 * @param {String} url
 * @return {Boolean}
 * @api public
 */

exports.isCrossDomain = function(url){
  url = exports.parse(url);
  return url.hostname !== location.hostname
    || url.port !== location.port
    || url.protocol !== location.protocol;
};

/**
 * Return default port for `protocol`.
 *
 * @param  {String} protocol
 * @return {String}
 * @api private
 */
function port (protocol){
  switch (protocol) {
    case 'http:':
      return 80;
    case 'https:':
      return 443;
    default:
      return location.port;
  }
}

});
require.register("segmentio-top-domain/index.js", function(exports, require, module){

var url = require('url');

// Official Grammar: http://tools.ietf.org/html/rfc883#page-56
// Look for tlds with up to 2-6 characters.

module.exports = function (urlStr) {

  var host     = url.parse(urlStr).hostname
    , topLevel = host.match(/[a-z0-9][a-z0-9\-]*[a-z0-9]\.[a-z\.]{2,6}$/i);

  return topLevel ? topLevel[0] : host;
};
});
require.register("timoxley-next-tick/index.js", function(exports, require, module){
"use strict"

if (typeof setImmediate == 'function') {
  module.exports = function(f){ setImmediate(f) }
}
// legacy node.js
else if (typeof process != 'undefined' && typeof process.nextTick == 'function') {
  module.exports = process.nextTick
}
// fallback for other environments / postMessage behaves badly on IE8
else if (typeof window == 'undefined' || window.ActiveXObject || !window.postMessage) {
  module.exports = function(f){ setTimeout(f) };
} else {
  var q = [];

  window.addEventListener('message', function(){
    var i = 0;
    while (i < q.length) {
      try { q[i++](); }
      catch (e) {
        q = q.slice(i);
        window.postMessage('tic!', '*');
        throw e;
      }
    }
    q.length = 0;
  }, true);

  module.exports = function(fn){
    if (!q.length) window.postMessage('tic!', '*');
    q.push(fn);
  }
}

});
require.register("skybox/lib/index.js", function(exports, require, module){

"use strict";
/*jslint browser: true, nomen: true*/

var Skybox = require('./skybox'),
    bind  = require('bind'),
    skybox = new Skybox();

skybox.autoinitialize();

module.exports = skybox;

bind(module.exports, module.exports.init);
bind(module.exports, module.exports.initialize);

});
require.register("skybox/lib/skybox.js", function(exports, require, module){

"use strict";
/*jslint browser: true, nomen: true, regexp: true*/

var Cookie      = require('./cookie'),
    Device      = require('./device'),
    User        = require('./user'),
    each        = require('each'),
    extend      = require('extend'),
    isEmpty     = require('is-empty'),
    querystring = require('querystring'),
    type        = require('type'),
    DEFAULTHOST = null,
    DEFAULTPORT = 80;

function Skybox() {
    this.apiKey = null;
    this.host = DEFAULTHOST;
    this.port = DEFAULTPORT;
    this.cookie = new Cookie();
    this.device = new Device();
    this.user = new User();
    this.initialized = false;
    this.resource(function () {
        return this.path().replace(/\/\d+(?=\/|\b)/g, "/:id");
    });
}

Skybox.prototype.initialize = function (apiKey, options) {
    var self = this;
    this.apiKey = apiKey;
    this.options(options);
    this.initialized = true;
};

Skybox.prototype.init = Skybox.prototype.inititialize;

/**
 * Attempts to automatically initialize based on data attributes in the script tag:
 * 
 * <script data-api-key="XXX-XXXX"/>
 */
Skybox.prototype.autoinitialize = function () {
    var arr, re, script = this.scriptElement();
    if (script === null) {
        return;
    }

    // Extract API key.
    if (this.apiKey === null && script.getAttribute("data-api-key") !== null) {
        this.initialize(script.getAttribute("data-api-key"));
    }

    // Extract host and port of script.
    if (this.host === null) {
        re = /^(?:https?:)?\/\/([^\/]+).*/;
        if (script.src.search(re) === 0) {
            arr = script.src.replace(re, "$1").split(":");
            this.host = arr[0];
            if (arr.length > 1 && !isNaN(parseInt(arr[1], 10))) {
                this.port = parseInt(arr[1], 10);
            }
        }
    }

    // Automatically track the page.
    this.page();
};

/**
 * Sets or retrieves the current options.
 */
Skybox.prototype.options = function (value) {
    if (arguments.length === 0) {
        return this._options;
    }
    var options = value || {};
    this._options = options;
    this.cookie.options(options.cookie);
    this.device.options(options.device);
    this.user.options(options.user);
};

/**
 * Identify a user by `id`.
 */
Skybox.prototype.identify = function (id) {
    this.user.identify(id);
    return this;
};

/**
 * Track an event that a user has triggered.
 */
Skybox.prototype.track = function (action) {
    var url, q, el, self = this,
        attr = this._options.mode === "test" ? "title" : "src",
        event = {
            channel: "web",
            resource: this.resource(),
            action: action,
            domain: this.domain(),
            path: this.path(),
        };

    // Ignore if not initialized yet.
    if (!this.initialized) {
        this.log("tracking not allowed before initialization");
        return this;
    }

    // Generate url.
    q = extend(event, {
        apiKey: this.apiKey,
        user: this.user.serialize(),
        device: this.device.serialize()
    });
    url = this.url("/track.png", q);

    // Send to server.
    el = document.createElement("img");
    el.width = el.height = 1;
    el[attr] = url;
    document.body.appendChild(el);

    // Remove the tracker image after it's had time to send.
    setTimeout(function () {
        try {
            document.body.removeChild(el);
        } catch (e) {
        }
    }, 100);

    return this;
};

/**
 * Tracks a page view. This is called automatically after initialization
 * but is useful to call for single-page apps.
 */
Skybox.prototype.page = function (name) {
    return this.track("view");
};

/**
 * Sets or retrieves the current resource. If set to a function then the
 * resource will be the result of the function.
 */
Skybox.prototype.resource = function (value) {
    if (arguments.length === 0) {
        return this._resource();
    }
    var v = (typeof (value) === "function" ? value : function () { return value; });
    this._resource = v;
};

/**
 * Retrieves the current domain of the page.
 */
Skybox.prototype.domain = function () {
    return window.location.host.replace(/:80$/, "");
};

/**
 * Retrieves the current path of the page.
 */
Skybox.prototype.path = function () {
    return window.location.pathname.replace(/\/+$/, "");
};

/**
 * Returns a URL with the appropriate host, port, path and query string.
 */
Skybox.prototype.url = function (path, q) {
    var i, key, params = {},
        str = "";

    // Setup scheme://host:port/path
    str += ('https:' === document.location.protocol ? "https://" : "http://");
    str += (isEmpty(this.host) ? "localhost" : this.host);
    str += (isEmpty(this.port) || this.port === 80 ? "" : ":" + this.port);
    str += path;

    // Flatten query parameters.
    if (type(q) === "object") {
        for (key in q) {
            if (q.hasOwnProperty(key)) {
                if (type(q[key]) === "object") {
                    for (i in q[key]) {
                        if (q[key].hasOwnProperty(i)) {
                            params[key + "." + i] = q[key][i];
                        }
                    }
                } else {
                    params[key] = q[key];
                }
            }
        }
    }

    // Append parameters to the end, if there are any.
    if (!isEmpty(params)) {
        str += "?" + querystring.stringify(params);
    }

    return str;
};

/**
 * Retrieves a reference to the first script element that loaded "skybox.js".
 */
Skybox.prototype.scriptElement = function () {
    var i, scripts = document.getElementsByTagName("script");
    for (i = 0; i < scripts.length; i += 1) {
        if (scripts[i].src.search(/\/skybox.js(?!\/)/) !== -1) {
            return scripts[i];
        }
    }
    return null;
};

Skybox.prototype.log = function (msg) {
    if (window.console) {
        window.console.log("[skybox.js]: " + msg);
    }
};

module.exports = Skybox;

Skybox.VERSION = Skybox.prototype.VERSION = '0.1.0';


});
require.register("skybox/lib/device.js", function(exports, require, module){

"use strict";
/*jslint browser: true, nomen: true*/

var Cookie = require('./cookie'),
    defaults = require('defaults');


/**
 * A Device represents the browser that the user is on. The device identifier
 * can be used while a user is unidentified.
 */
function Device(options) {
    this.cookie = new Cookie();
    this.options(options);
    this.initialize();
}

/**
 * Initializes the device with a new identifier if it doesn't have one yet.
 */
Device.prototype.initialize = function () {
    if (this.id() === null) {
        var id = "xxxxxxxxxxxx4xxxyxxxxxxxxxxxxxxx".replace(/[xy]/g, function (c) {
            /*jslint bitwise: true*/
            var r = (Math.random() * 16) | 0,
                v = c === 'x' ? r : (r & 0x3 | 0x8);
            /*jslint bitwise: false*/
            return v.toString(16);
        });
        this.cookie.set(this._options.cookie.key, id);
    }
};

/**
 * Sets or retrieves the options on the device.
 */
Device.prototype.options = function (value) {
    if (arguments.length === 0) {
        return this._options;
    }
    var options = value || {};
    defaults(options, {cookie: {key: 'trackjs_DeviceID'}});
    this._options = options;
};

/**
 * Retrieve's the device identifier.
 */
Device.prototype.id = function (id) {
    if (this._options.mode === "test") {
        return "x";
    }
    return this.cookie.get(this._options.cookie.key);
};

/**
 * Serializes the Device into a hash.
 */
Device.prototype.serialize = function () {
    return {id: this.id()};
};

module.exports = Device;


});
require.register("skybox/lib/cookie.js", function(exports, require, module){

"use strict";
/*jslint browser: true, nomen: true*/

var bind = require('bind'),
    cookie = require('cookie'),
    clone = require('clone'),
    defaults = require('defaults'),
    topDomain = require('top-domain');


function Cookie(options) {
    this.options(options);
}

/**
 * Get or set the cookie options
 */
Cookie.prototype.options = function (value) {
    if (arguments.length === 0) {
        return this._options;
    }

    var options = value || {},
        domain = '.' + topDomain(window.location.href);

    if (domain === '.localhost') {
        domain = '';
    }

    defaults(options, {
        maxage  : 31536000000, // default to a year
        path    : '/',
        domain  : domain
    });

    this._options = options;
};


/**
 * Set a value in our cookie
 */
Cookie.prototype.set = function (key, value) {
    try {
        value = JSON.stringify(value);
        cookie(key, value, clone(this._options));
        return true;
    } catch (e) {
        return false;
    }
};


/**
 * Get a value from our cookie.s
 */
Cookie.prototype.get = function (key) {
    try {
        var value = cookie(key);
        value = value ? JSON.parse(value) : null;
        return value;
    } catch (e) {
        return null;
    }
};


/**
 * Remove a value from the cookie.
 */
Cookie.prototype.remove = function (key) {
    try {
        cookie(key, null, clone(this._options));
        return true;
    } catch (e) {
        return false;
    }
};

module.exports = Cookie;

});
require.register("skybox/lib/user.js", function(exports, require, module){

"use strict";
/*jslint browser: true, nomen: true*/

var clone = require('clone'),
    Cookie = require('./cookie'),
    clone = require('clone'),
    defaults = require('defaults'),
    extend = require('extend'),
    isEmpty = require('is-empty');


function User(options) {
    this.cookie = new Cookie();
    this.options(options);
    this.id(null);
}

User.prototype.options = function (value) {
    if (arguments.length === 0) {
        return this._options;
    }

    var options = value || {};
    defaults(options, {
        cookie: {
            key: 'trackjs_UserID',
        },
    });

    this._options = options;
};

/**
 * Get or set the user's `id`.
 */
User.prototype.id = function (id) {
    if (arguments.length === 0) {
        return this.cookie.get(this._options.cookie.key);
    }
    this.cookie.set(this._options.cookie.key, id);
};

/**
 * Identity the user with an `id`.
 */
User.prototype.identify = function (id) {
    this.id(id);
};

/**
 * Log the user out, resetting `id` to blank.
 */
User.prototype.logout = function () {
    this.id(null);
    this.cookie.remove(this._options.cookie.key);
};

/**
 * Serializes the User into a hash.
 */
User.prototype.serialize = function () {
    var obj = {};
    if (this.id()) {
        obj.id = this.id();
    }
    return obj;
};


module.exports = User;

});











require.alias("avetisk-defaults/index.js", "skybox/deps/defaults/index.js");
require.alias("avetisk-defaults/index.js", "defaults/index.js");

require.alias("component-bind/index.js", "skybox/deps/bind/index.js");
require.alias("component-bind/index.js", "bind/index.js");

require.alias("component-clone/index.js", "skybox/deps/clone/index.js");
require.alias("component-clone/index.js", "clone/index.js");
require.alias("component-type/index.js", "component-clone/deps/type/index.js");

require.alias("component-cookie/index.js", "skybox/deps/cookie/index.js");
require.alias("component-cookie/index.js", "cookie/index.js");

require.alias("component-each/index.js", "skybox/deps/each/index.js");
require.alias("component-each/index.js", "each/index.js");
require.alias("component-to-function/index.js", "component-each/deps/to-function/index.js");
require.alias("component-props/index.js", "component-to-function/deps/props/index.js");

require.alias("component-type/index.js", "component-each/deps/type/index.js");

require.alias("component-querystring/index.js", "skybox/deps/querystring/index.js");
require.alias("component-querystring/index.js", "querystring/index.js");
require.alias("component-trim/index.js", "component-querystring/deps/trim/index.js");

require.alias("component-type/index.js", "skybox/deps/type/index.js");
require.alias("component-type/index.js", "type/index.js");

require.alias("ianstormtaylor-is-empty/index.js", "skybox/deps/is-empty/index.js");
require.alias("ianstormtaylor-is-empty/index.js", "is-empty/index.js");

require.alias("segmentio-extend/index.js", "skybox/deps/extend/index.js");
require.alias("segmentio-extend/index.js", "extend/index.js");

require.alias("segmentio-top-domain/index.js", "skybox/deps/top-domain/index.js");
require.alias("segmentio-top-domain/index.js", "skybox/deps/top-domain/index.js");
require.alias("segmentio-top-domain/index.js", "top-domain/index.js");
require.alias("component-url/index.js", "segmentio-top-domain/deps/url/index.js");

require.alias("segmentio-top-domain/index.js", "segmentio-top-domain/index.js");
require.alias("timoxley-next-tick/index.js", "skybox/deps/next-tick/index.js");
require.alias("timoxley-next-tick/index.js", "next-tick/index.js");

require.alias("skybox/lib/index.js", "skybox/index.js");if (typeof exports == "object") {
  module.exports = require("skybox");
} else if (typeof define == "function" && define.amd) {
  define(function(){ return require("skybox"); });
} else {
  this["skybox"] = require("skybox");
}})();