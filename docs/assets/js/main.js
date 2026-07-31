'use strict';

document.addEventListener('DOMContentLoaded', function () {
  const SVG_NS = 'http://www.w3.org/2000/svg';

  function createBaseSvg() {
    const svg = document.createElementNS(SVG_NS, 'svg');
    svg.setAttribute('width', '15');
    svg.setAttribute('height', '15');
    svg.setAttribute('viewBox', '0 0 24 24');
    svg.setAttribute('fill', 'none');
    svg.setAttribute('stroke', 'currentColor');
    svg.setAttribute('stroke-width', '2');
    svg.setAttribute('stroke-linecap', 'round');
    svg.setAttribute('stroke-linejoin', 'round');
    return svg;
  }

  function createCopyIcon() {
    const svg = createBaseSvg();

    const path = document.createElementNS(SVG_NS, 'path');
    path.setAttribute('d', 'M16 4h2a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h2');

    const rect = document.createElementNS(SVG_NS, 'rect');
    rect.setAttribute('x', '8');
    rect.setAttribute('y', '2');
    rect.setAttribute('width', '8');
    rect.setAttribute('height', '4');
    rect.setAttribute('rx', '1');
    rect.setAttribute('ry', '1');

    svg.appendChild(path);
    svg.appendChild(rect);
    return svg;
  }

  function createCheckIcon() {
    const svg = createBaseSvg();

    const polyline = document.createElementNS(SVG_NS, 'polyline');
    polyline.setAttribute('points', '20 6 9 17 4 12');

    svg.appendChild(polyline);
    return svg;
  }

  function fallbackCopyToClipboard(text) {
    return new Promise(function (resolve, reject) {
      try {
        const textArea = document.createElement('textarea');
        textArea.value = text;
        textArea.style.position = 'fixed';
        textArea.style.left = '-9999px';
        textArea.style.top = '-9999px';
        document.body.appendChild(textArea);
        textArea.focus();
        textArea.select();
        const successful = document.execCommand('copy');
        document.body.removeChild(textArea);
        if (successful) {
          resolve();
        } else {
          reject(new Error('Fallback copy failed'));
        }
      } catch (err) {
        reject(err);
      }
    });
  }

  function copyToClipboard(text) {
    if (navigator && navigator.clipboard && typeof navigator.clipboard.writeText === 'function') {
      return navigator.clipboard.writeText(text).catch(function () {
        return fallbackCopyToClipboard(text);
      });
    }
    return fallbackCopyToClipboard(text);
  }

  const blocks = document.querySelectorAll('div.highlighter-rouge');
  blocks.forEach(function (block) {
    if (block.querySelector('.copy-btn')) return;

    const btn = document.createElement('button');
    btn.className = 'copy-btn';
    btn.type = 'button';
    btn.setAttribute('aria-label', 'Copy code to clipboard');
    btn.setAttribute('title', 'Copy');
    btn.appendChild(createCopyIcon());

    btn.addEventListener('click', function () {
      const codeEl = block.querySelector('code') || block.querySelector('pre');
      let text = '';
      if (codeEl) {
        text = codeEl.textContent || '';
      } else {
        const clone = block.cloneNode(true);
        const existingBtn = clone.querySelector('.copy-btn');
        if (existingBtn) existingBtn.remove();
        text = clone.textContent || '';
      }

      copyToClipboard(text)
        .then(function () {
          if (btn._resetTimer) {
            clearTimeout(btn._resetTimer);
          }
          btn.replaceChildren(createCheckIcon());
          btn.setAttribute('title', 'Copied!');
          btn.setAttribute('aria-label', 'Copied!');
          btn.classList.add('copied');
          btn._resetTimer = setTimeout(function () {
            btn.replaceChildren(createCopyIcon());
            btn.setAttribute('title', 'Copy');
            btn.setAttribute('aria-label', 'Copy code to clipboard');
            btn.classList.remove('copied');
            btn._resetTimer = null;
          }, 2000);
        })
        .catch(function () {
          // ignore copy failures silently
        });
    });

    block.appendChild(btn);
  });
});


