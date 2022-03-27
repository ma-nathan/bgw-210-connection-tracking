FROM chromedp/headless-shell:92.0.4515.20

RUN apt -y update ; echo ignore error status return
RUN apt -y install procps

ADD files/entrypoint.sh /
ADD settings.ini /
ADD files/bgw210 /

ENTRYPOINT ["/entrypoint.sh"]


