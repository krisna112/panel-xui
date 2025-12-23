[English](/README.md) | [ÙØ§Ø±Ø³ÛŒ](/README.fa_IR.md) | [Ø§Ù„Ø¹Ø±Ø¨ÙŠØ©](/README.ar_EG.md) |  [ä¸­æ–‡](/README.zh_CN.md) | [EspaÃ±ol](/README.es_ES.md) | [Ð ÑƒÑÑÐºÐ¸Ð¹](/README.ru_RU.md)

<p align="center">
  <picture>
    <source media="(prefers-color-scheme: dark)" srcset="./media/3x-ui-dark.png">
    <img alt="3x-ui" src="./media/3x-ui-light.png">
  </picture>
</p>

[![Release](https://img.shields.io/github/v/release/krisna112/panel-xui.svg)](https://github.com/krisna112/panel-xui/releases)
[![Build](https://img.shields.io/github/actions/workflow/status/krisna112/panel-xui/release.yml.svg)](https://github.com/krisna112/panel-xui/actions)
[![GO Version](https://img.shields.io/github/go-mod/go-version/krisna112/panel-xui.svg)](#)
[![Downloads](https://img.shields.io/github/downloads/krisna112/panel-xui/total.svg)](https://github.com/krisna112/panel-xui/releases/latest)
[![License](https://img.shields.io/badge/license-GPL%20V3-blue.svg?longCache=true)](https://www.gnu.org/licenses/gpl-3.0.en.html)
[![Go Reference](https://pkg.go.dev/badge/github.com/krisna112/panel-xui/v2.svg)](https://pkg.go.dev/github.com/krisna112/panel-xui/v2)
[![Go Report Card](https://goreportcard.com/badge/github.com/krisna112/panel-xui/v2)](https://goreportcard.com/report/github.com/krisna112/panel-xui/v2)

**3X-UI** â€” ÛŒÚ© Ù¾Ù†Ù„ Ú©Ù†ØªØ±Ù„ Ù¾ÛŒØ´Ø±ÙØªÙ‡ Ù…Ø¨ØªÙ†ÛŒ Ø¨Ø± ÙˆØ¨ Ø¨Ø§ Ú©Ø¯ Ø¨Ø§Ø² Ú©Ù‡ Ø¨Ø±Ø§ÛŒ Ù…Ø¯ÛŒØ±ÛŒØª Ø³Ø±ÙˆØ± Xray-core Ø·Ø±Ø§Ø­ÛŒ Ø´Ø¯Ù‡ Ø§Ø³Øª. Ø§ÛŒÙ† Ù¾Ù†Ù„ ÛŒÚ© Ø±Ø§Ø¨Ø· Ú©Ø§Ø±Ø¨Ø±ÛŒ Ø¢Ø³Ø§Ù† Ø¨Ø±Ø§ÛŒ Ù¾ÛŒÚ©Ø±Ø¨Ù†Ø¯ÛŒ Ùˆ Ù†Ø¸Ø§Ø±Øª Ø¨Ø± Ù¾Ø±ÙˆØªÚ©Ù„â€ŒÙ‡Ø§ÛŒ Ù…Ø®ØªÙ„Ù VPN Ùˆ Ù¾Ø±Ø§Ú©Ø³ÛŒ Ø§Ø±Ø§Ø¦Ù‡ Ù…ÛŒâ€ŒØ¯Ù‡Ø¯.

> [!IMPORTANT]
> Ø§ÛŒÙ† Ù¾Ø±ÙˆÚ˜Ù‡ ÙÙ‚Ø· Ø¨Ø±Ø§ÛŒ Ø§Ø³ØªÙØ§Ø¯Ù‡ Ø´Ø®ØµÛŒ Ùˆ Ø§Ø±ØªØ¨Ø§Ø·Ø§Øª Ø§Ø³ØªØŒ Ù„Ø·ÙØ§Ù‹ Ø§Ø² Ø¢Ù† Ø¨Ø±Ø§ÛŒ Ø§Ù‡Ø¯Ø§Ù ØºÛŒØ±Ù‚Ø§Ù†ÙˆÙ†ÛŒ Ø§Ø³ØªÙØ§Ø¯Ù‡ Ù†Ú©Ù†ÛŒØ¯ØŒ Ù„Ø·ÙØ§Ù‹ Ø§Ø² Ø¢Ù† Ø¯Ø± Ù…Ø­ÛŒØ· ØªÙˆÙ„ÛŒØ¯ Ø§Ø³ØªÙØ§Ø¯Ù‡ Ù†Ú©Ù†ÛŒØ¯.

Ø¨Ù‡ Ø¹Ù†ÙˆØ§Ù† ÛŒÚ© Ù†Ø³Ø®Ù‡ Ø¨Ù‡Ø¨ÙˆØ¯ ÛŒØ§ÙØªÙ‡ Ø§Ø² Ù¾Ø±ÙˆÚ˜Ù‡ Ø§ØµÙ„ÛŒ X-UIØŒ 3X-UI Ù¾Ø§ÛŒØ¯Ø§Ø±ÛŒ Ø¨Ù‡ØªØ±ØŒ Ù¾Ø´ØªÛŒØ¨Ø§Ù†ÛŒ Ú¯Ø³ØªØ±Ø¯Ù‡â€ŒØªØ± Ø§Ø² Ù¾Ø±ÙˆØªÚ©Ù„â€ŒÙ‡Ø§ Ùˆ ÙˆÛŒÚ˜Ú¯ÛŒâ€ŒÙ‡Ø§ÛŒ Ø§Ø¶Ø§ÙÛŒ Ø±Ø§ Ø§Ø±Ø§Ø¦Ù‡ Ù…ÛŒâ€ŒØ¯Ù‡Ø¯.

## Ø´Ø±ÙˆØ¹ Ø³Ø±ÛŒØ¹

```
bash <(curl -Ls https://raw.githubusercontent.com/krisna112/panel-xui/main/install.sh)
```

Ø¨Ø±Ø§ÛŒ Ù…Ø³ØªÙ†Ø¯Ø§Øª Ú©Ø§Ù…Ù„ØŒ Ù„Ø·ÙØ§Ù‹ Ø¨Ù‡ [ÙˆÛŒÚ©ÛŒ Ù¾Ø±ÙˆÚ˜Ù‡](https://github.com/krisna112/panel-xui/wiki) Ù…Ø±Ø§Ø¬Ø¹Ù‡ Ú©Ù†ÛŒØ¯.

## ØªØ´Ú©Ø± ÙˆÛŒÚ˜Ù‡ Ø§Ø²

- [alireza0](https://github.com/alireza0/)

## Ù‚Ø¯Ø±Ø¯Ø§Ù†ÛŒ

- [Iran v2ray rules](https://github.com/chocolate4u/Iran-v2ray-rules) (Ù…Ø¬ÙˆØ²: **GPL-3.0**): _Ù‚ÙˆØ§Ù†ÛŒÙ† Ù…Ø³ÛŒØ±ÛŒØ§Ø¨ÛŒ Ø¨Ù‡Ø¨ÙˆØ¯ ÛŒØ§ÙØªÙ‡ v2ray/xray Ùˆ v2ray/xray-clients Ø¨Ø§ Ø¯Ø§Ù…Ù†Ù‡â€ŒÙ‡Ø§ÛŒ Ø§ÛŒØ±Ø§Ù†ÛŒ Ø¯Ø§Ø®Ù„ÛŒ Ùˆ ØªÙ…Ø±Ú©Ø² Ø¨Ø± Ø§Ù…Ù†ÛŒØª Ùˆ Ù…Ø³Ø¯ÙˆØ¯ Ú©Ø±Ø¯Ù† ØªØ¨Ù„ÛŒØºØ§Øª._
- [Russia v2ray rules](https://github.com/runetfreedom/russia-v2ray-rules-dat) (Ù…Ø¬ÙˆØ²: **GPL-3.0**): _Ø§ÛŒÙ† Ù…Ø®Ø²Ù† Ø´Ø§Ù…Ù„ Ù‚ÙˆØ§Ù†ÛŒÙ† Ù…Ø³ÛŒØ±ÛŒØ§Ø¨ÛŒ V2Ray Ø¨Ù‡â€ŒØ±ÙˆØ²Ø±Ø³Ø§Ù†ÛŒ Ø´Ø¯Ù‡ Ø®ÙˆØ¯Ú©Ø§Ø± Ø¨Ø± Ø§Ø³Ø§Ø³ Ø¯Ø§Ø¯Ù‡â€ŒÙ‡Ø§ÛŒ Ø¯Ø§Ù…Ù†Ù‡â€ŒÙ‡Ø§ Ùˆ Ø¢Ø¯Ø±Ø³â€ŒÙ‡Ø§ÛŒ Ù…Ø³Ø¯ÙˆØ¯ Ø´Ø¯Ù‡ Ø¯Ø± Ø±ÙˆØ³ÛŒÙ‡ Ø§Ø³Øª._

## Ù¾Ø´ØªÛŒØ¨Ø§Ù†ÛŒ Ø§Ø² Ù¾Ø±ÙˆÚ˜Ù‡

**Ø§Ú¯Ø± Ø§ÛŒÙ† Ù¾Ø±ÙˆÚ˜Ù‡ Ø¨Ø±Ø§ÛŒ Ø´Ù…Ø§ Ù…ÙÛŒØ¯ Ø§Ø³ØªØŒ Ù…ÛŒâ€ŒØªÙˆØ§Ù†ÛŒØ¯ Ø¨Ù‡ Ø¢Ù† ÛŒÚ©**:star2: Ø¨Ø¯Ù‡ÛŒØ¯

<a href="https://www.buymeacoffee.com/MHSanaei" target="_blank">
<img src="./media/default-yellow.png" alt="Buy Me A Coffee" style="height: 70px !important;width: 277px !important;" >
</a>

</br>
<a href="https://nowpayments.io/donation/hsanaei" target="_blank" rel="noreferrer noopener">
   <img src="./media/donation-button-black.svg" alt="Crypto donation button by NOWPayments">
</a>

## Ø³ØªØ§Ø±Ù‡â€ŒÙ‡Ø§ Ø¯Ø± Ø·ÙˆÙ„ Ø²Ù…Ø§Ù†

[![Stargazers over time](https://starchart.cc/krisna112/panel-xui.svg?variant=adaptive)](https://starchart.cc/krisna112/panel-xui) 
