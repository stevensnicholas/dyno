FROM  mcr.microsoft.com/restlerfuzzer/restler:v7.4.0
COPY rest_interface.py /
COPY requirements.txt /
RUN pip install -r requirements.txt
WORKDIR /
EXPOSE 5000
ENTRYPOINT ["python3","rest_interface.py"]