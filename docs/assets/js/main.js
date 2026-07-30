document.addEventListener('DOMContentLoaded', function () {
  var iconCopy = '<svg xmlns="http://www.w3.org/2000/svg" width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M16 4h2a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h2"></path><rect x="8" y="2" width="8" height="4" rx="1" ry="1"></rect></svg>';
  var iconCheck = '<svg xmlns="http://www.w3.org/2000/svg" width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polyline points="20 6 9 17 4 12"></polyline></svg>';

  function copyToClipboard(text) {
    if (navigator && navigator.clipboard && typeof navigator.clipboard.writeText === 'function') {
      return navigator.clipboard.writeText(text);
    }
    return new Promise(function (resolve, reject) {
      try {
        var textArea = document.createElement('textarea');
        textArea.value = text;
        textArea.style.position = 'fixed';
        textArea.style.left = '-9999px';
        textArea.style.top = '-9999px';
        document.body.appendChild(textArea);
        textArea.focus();
        textArea.select();
        var successful = document.execCommand('copy');
        document.body.removeChild(textArea);
        if (successful) {
          resolve();
        } else {
          reject(new Error('execCommand copy failed'));
        }
      } catch (err) {
        reject(err);
      }
    });
  }

  var blocks = document.querySelectorAll('div.highlighter-rouge');
  blocks.forEach(function (block) {
    if (block.querySelector('.copy-btn')) return;
    var btn = document.createElement('button');
    btn.className = 'copy-btn';
    btn.type = 'button';
    btn.setAttribute('aria-label', 'Copy code to clipboard');
    btn.setAttribute('title', 'Copy');
    btn.innerHTML = iconCopy;

    btn.addEventListener('click', function () {
      var codeEl = block.querySelector('code') || block.querySelector('pre');
      var text = '';
      if (codeEl) {
        text = codeEl.textContent;
      } else {
        var clone = block.cloneNode(true);
        var existingBtn = clone.querySelector('.copy-btn');
        if (existingBtn) existingBtn.remove();
        text = clone.textContent;
      }

      copyToClipboard(text).then(function () {
        btn.innerHTML = iconCheck;
        btn.setAttribute('title', 'Copied!');
        btn.setAttribute('aria-label', 'Copied!');
        btn.classList.add('copied');
        setTimeout(function () {
          btn.innerHTML = iconCopy;
          btn.setAttribute('title', 'Copy');
          btn.setAttribute('aria-label', 'Copy code to clipboard');
          btn.classList.remove('copied');
        }, 2000);
      }).catch(function (err) {
        console.error('Failed to copy code: ', err);
      });
    });

    block.appendChild(btn);
  });
});
