#!/usr/bin/env python
import sqlalchemy as sa
from sqlalchemy.orm import sessionmaker
from sqlalchemy.ext.declarative import declarative_base
from sklearn.datasets import load_iris
import pandas as pd

URL = 'mysql+pymysql://root:@0.0.0.0/test_db?charset=utf8'
Base = declarative_base()


class Iris(Base):
    __tablename__ = 'iris'
    id = sa.Column(sa.Integer, primary_key=True)
    sepal_length = sa.Column(sa.Float)
    sepal_width = sa.Column(sa.Float)
    petal_length = sa.Column(sa.Float)
    petal_width = sa.Column(sa.Float)
    label = sa.Column(sa.Integer)

    def __init__(self, sepal_length, sepal_width, petal_length, petal_width, label):
        self.sepal_length = sepal_length
        self.sepal_width = sepal_width
        self.petal_length = petal_length
        self.petal_width = petal_width
        self.label = label


def main():
    # connect DB
    engine = sa.create_engine(URL, echo=True)

    # create table
    Base.metadata.create_all(engine)

    # create session
    Session = sessionmaker(bind=engine)
    session = Session()
    dummy_data = create_dummy_data()

    # insert dummy data
    session.add_all(dummy_data)
    session.commit()

    # update rows
    update_row_id = [i for i in range(1, 150) if i % 2 == 1]
    for u_id in update_row_id:
        items = session.query(Iris).filter(Iris.id==u_id).first()
        items.label = 3
        session.commit()

    # delete rows
    delete_row_id = [i for i in range(1, 150) if i % 2 == 0]
    for d_id in delete_row_id:
        items = session.query(Iris).filter(Iris.id==d_id).delete()
        session.commit()


def create_dummy_data():
    data = load_iris()
    raw_data = pd.DataFrame(data.get('data'), columns=data.get('feature_names'))
    labels = pd.DataFrame(data.get('target'), columns=['label'])
    raw_data = pd.concat([raw_data, labels], axis=1)
    raw_data.head()

    alchemy_data = [Iris(sepal_length=float(raw[0]),
                         sepal_width=float(raw[1]),
                         petal_length=float(raw[2]),
                         petal_width=float(raw[3]),
                         label=int(raw[4])) for raw in raw_data.values]
    return alchemy_data


if __name__ == "__main__":
    main()