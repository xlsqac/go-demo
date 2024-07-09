def main():
    try:
        1 /0
    finally:
        print(12)


if __name__ == "__main__":
    print(main())
    print(13)
