FROM chromedp/headless-shell:92.0.4515.20

ADD files/entrypoint.sh /
ADD settings.ini /
ADD files/bgw210 /

ENTRYPOINT ["/entrypoint.sh"]


